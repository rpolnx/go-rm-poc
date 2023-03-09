package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"time"
)

func NewPgClient(cfgApp *config.App, cfgDb *config.Db) (*pg.DB, error) {
	initialDbTime := time.Now()

	logrus.Infof("[Postgres DB] Initializing pg dependency")

	db := pg.Connect(&pg.Options{
		Addr:            fmt.Sprintf("%s:%d", cfgDb.Host, cfgDb.Port),
		User:            cfgDb.Username,
		Password:        cfgDb.Password,
		Database:        cfgDb.DbName,
		ApplicationName: cfgApp.Name,
		PoolSize:        10,
		ReadTimeout:     time.Duration(cfgDb.Timeout) * time.Second,
		WriteTimeout:    time.Duration(cfgDb.Timeout) * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfgDb.Timeout)*time.Second)
	defer cancel()

	err := db.Ping(ctx)

	deltaDb := time.Since(initialDbTime).Milliseconds()
	logrus.Infof("[Postgres DB] Finalized pg dependency in %dms", deltaDb)

	if err != nil {
		return nil, err
	}

	initialDbMigrationsTime := time.Now()
	logrus.Infof("[Postgres DB] Initializing pg migrations")

	if err = RunMigrations(db, cfgDb); err != nil {
		return nil, err
	}

	deltaMigrations := time.Since(initialDbMigrationsTime).Milliseconds()
	logrus.Infof("[Postgres DB] Finalized pg migrations in %dms", deltaMigrations)

	return db, nil
}

func RunMigrations(dbClient *pg.DB, cfg *config.Db) error {
	c := migrations.NewCollection()
	c.DisableSQLAutodiscover(true)
	err := c.DiscoverSQLMigrations(cfg.MigrationPath)
	if err != nil {
		logrus.Errorf("[Postgres DB] setting DiscoverSQLMigrations() to path %s failed: %s", cfg.MigrationPath, err.Error())
		return err
	}

	c.SetTableName(fmt.Sprintf("%s.%s", cfg.Schema, cfg.MigrationTable))

	oldVersion, newVersion, err := c.Run(dbClient, "up")

	if err == nil {
		logrus.Debugf("[Postgres DB] version is %d\n", newVersion)
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"path":        cfg.MigrationPath,
		"oldVersion":  oldVersion,
		"newVersion":  newVersion,
		"err.Error()": err.Error(),
	}).Debugf("[Postgres DB] Initializing migrations mechanism")

	//init migrations mechanism for the 1st run
	if newVersion == 0 {
		_, _, err = c.Run(dbClient, "init")
		if err != nil {
			return err
		}

		oldVersion, newVersion, err = c.Run(dbClient, "up")
	}

	if err != nil { //this check is because it can be changed by new c.Run(db) two lines upper
		return err
	}

	if newVersion != oldVersion {
		logrus.Infof("[Postgres DB] migrated from version %d to %d", oldVersion, newVersion)
	} else {
		logrus.Debugf("[Postgres DB] version is %d\n", oldVersion)
	}

	return nil
}
