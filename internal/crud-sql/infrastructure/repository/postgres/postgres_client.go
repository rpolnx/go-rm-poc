package postgres

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"time"
)

func NewPgClient(cfg config.Db) (*gorm.DB, error) {
	initialDbTime := time.Now()

	logrus.Infof("[Postgres DB] Initializing mongo client dependencies")

	//opts := options.Client().ApplyURI(mongoC.Uri)
	opts := &gorm.Config{
		//Logger: logrus.Logger{}
		//ConnPool: nil
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.DbName, cfg.Port, cfg.SslMode, cfg.Timezone)
	client, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
		DriverName:           cfg.DriverName,
	}), opts)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	client = client.WithContext(ctx)

	//err = client.Ping(ctx, readpref.Primary())
	//
	//if err != nil {
	//	logrus.Errorf("[Postgres DB] Error initializing mongo client dependencies")
	//	return nil, err
	//}

	deltaDb := time.Since(initialDbTime).Milliseconds()
	logrus.Infof("[Postgres DB] Finalized store dependency in %dms", deltaDb)

	sqlDB, err := client.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return client, nil
}
