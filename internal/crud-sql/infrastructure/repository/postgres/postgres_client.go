package postgres

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/migrate"
	"golang.org/x/net/context"
	"os"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"time"

	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewPgClient(cfgApp *config.App, cfgDb *config.Db) (*bun.DB, error) {
	initialDbTime := time.Now()

	logrus.Infof("[Postgres DB] Initializing pg dependency")

	pgCon := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(fmt.Sprintf("%s:%d", cfgDb.Host, cfgDb.Port)),
		pgdriver.WithInsecure(true), // disable TLS
		//pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(cfgDb.Username),
		pgdriver.WithPassword(cfgDb.Password),
		pgdriver.WithDatabase(cfgDb.DbName),
		pgdriver.WithApplicationName(cfgApp.Name),
		pgdriver.WithTimeout(time.Duration(cfgDb.Timeout)*time.Second),
		pgdriver.WithDialTimeout(time.Duration(cfgDb.Timeout)*time.Second),
		pgdriver.WithReadTimeout(time.Duration(cfgDb.Timeout)*time.Second),
		pgdriver.WithWriteTimeout(time.Duration(cfgDb.Timeout)*time.Second),
		pgdriver.WithConnParams(map[string]interface{}{
			"search_path": cfgDb.Schema,
		}),
	)

	sqldb := sql.OpenDB(pgCon)

	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	err := db.Ping()

	if err != nil {
		return nil, err
	}

	deltaDb := time.Since(initialDbTime).Milliseconds()
	logrus.Infof("[Postgres DB] Finalized pg dependency in %dms", deltaDb)

	initialDbMigrationsTime := time.Now()
	logrus.Infof("[Postgres DB] Initializing pg migrations")

	if err = RunMigrations(db, cfgDb); err != nil {
		return nil, err
	}

	deltaMigrations := time.Since(initialDbMigrationsTime).Milliseconds()
	logrus.Infof("[Postgres DB] Finalized pg migrations in %dms", deltaMigrations)

	return db, nil
}

func RunMigrations(dbClient *bun.DB, cfg *config.Db) error {

	dirFs := os.DirFS("configs/migrations")

	migrations := migrate.NewMigrations()

	if err := migrations.Discover(dirFs); err != nil {
		logrus.Errorf("[Postgres DB] setting DiscoverCaller() to path %s failed: %s", cfg.MigrationPath, err.Error())
		return err
	}

	tableName := migrate.WithTableName(fmt.Sprintf("%s.%s", cfg.Schema, cfg.MigrationTable))

	migrator := migrate.NewMigrator(dbClient, migrations, tableName)

	timeoutCtx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancelCtx()

	if err := migrator.Init(timeoutCtx); err != nil {
		logrus.Errorf("[Postgres DB] error setting up first version %v\n", err.Error())
		return err
	}

	group, err := migrator.Migrate(timeoutCtx)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"path":        cfg.MigrationPath,
			"group":       group,
			"err.Error()": err.Error(),
		}).Errorf("[Postgres DB] Finishing migrations mechanism")

		return err
	}

	return nil
}
