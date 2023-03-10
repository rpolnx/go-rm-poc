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

	companyRepository := repository.NewCompanyRepository(dbClient)
	userRepository := repository.NewUserRepository(dbClient)

	companyService := service.NewCompanyService(companyRepository)
	userService := service.NewUserService(userRepository)

	healthcheckController := controller.NewHealthcheckController(cfg)
	userController := controller.NewUserController(cfg, userService)
	companyController := controller.NewCompanyController(cfg, companyService)

	routes.NewHealthcheckRoute(server, healthcheckController)
	routes.NewUserRoute(server, userController)
	routes.NewCompanyRoute(server, companyController)

	return server, nil
}
