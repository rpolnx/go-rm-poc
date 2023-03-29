package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/sirupsen/logrus"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/controller"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/routes"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/service"
	"rpolnx.com.br/crud-sql/internal/crud-sql/infrastructure/repository"
	"rpolnx.com.br/crud-sql/internal/crud-sql/infrastructure/repository/postgres"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

func LoadServer(cfg *config.Configuration) (*gin.Engine, error) {
	logrus.Println("Initializing dependencies")

	ctx := context.Background()

	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, closer := jaeger.NewTracer(
		cfg.App.Name,
		jaeger.NewConstSampler(true),
		jaeger.NewInMemoryReporter(),
		jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, propagator),
		jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, propagator),
		jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
	)
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	server := gin.Default()

	server.Use(ginhttp.Middleware(tracer))

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
