package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/controller"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/routes"
)

func LoadServer(cfg *config.Configuration) (*gin.Engine, error) {
	log.Println("Initializing dependencies")

	server := gin.Default()

	envController := controller.NewEnvController(cfg)
	healthcheckController := controller.NewHealthcheckController(cfg)
	userController := controller.NewUserController(cfg)

	routes.NewEnvRoute(server, envController)
	routes.NewHealthcheckRoute(server, healthcheckController)
	routes.NewUserRoute(server, userController)

	return server, nil
}
