package config

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type DbLogger struct{}

func (d DbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d DbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, err := q.FormattedQuery()
	logrus.Debugln(fmt.Sprintf(string(query)))
	return err
}

func init() {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
