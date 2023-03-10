package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/server"
)

func main() {
	logrus.Infof("Initializing server")

	loadConfig, err := config.LoadConfig()

	if err != nil {
		logrus.Fatal(err)
	}

	currentServer, err := server.LoadServer(loadConfig)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Starting server on port %d", loadConfig.App.Port)

	if err = currentServer.Run(fmt.Sprintf("localhost:%d", loadConfig.App.Port)); err != nil {
		logrus.Fatal(err)
	}
}
