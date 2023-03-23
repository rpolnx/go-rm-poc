package config

import (
	"github.com/sirupsen/logrus"
)

func init() {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
