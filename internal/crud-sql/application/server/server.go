package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/controller"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/routes"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/service"
	"rpolnx.com.br/crud-sql/internal/crud-sql/infrastructure/repository"
	"rpolnx.com.br/crud-sql/internal/crud-sql/infrastructure/repository/postgres"
)

func LoadServer(cfg *config.Configuration) (*gin.Engine, error) {
	logrus.Println("Initializing dependencies")

	server := gin.Default()

	dbClient, err := postgres.NewPgClient(&cfg.App, &cfg.Db)

	if err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(dbClient)

	userService := service.NewUserService(userRepository)

	envController := controller.NewEnvController(cfg)
	healthcheckController := controller.NewHealthcheckController(cfg)
	userController := controller.NewUserController(cfg, userService)

	routes.NewEnvRoute(server, envController)
	routes.NewHealthcheckRoute(server, healthcheckController)
	routes.NewUserRoute(server, userController)

	return server, nil
}
