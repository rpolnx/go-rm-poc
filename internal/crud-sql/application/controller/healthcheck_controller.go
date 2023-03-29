package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
)

type HealthcheckController interface {
	GetServerStatus(c *gin.Context)
}

type healthcheckController struct {
	cfg *config.Configuration
}

func (h *healthcheckController) GetServerStatus(c *gin.Context) {
	m := map[string]interface{}{"status": "OK"}

	span := opentracing.SpanFromContext(c.Request.Context())

	context := span.Context()
	fmt.Println(context)

	context

	logrus.Info("Example")

	c.JSON(http.StatusOK, m)
}

func NewHealthcheckController(cfg *config.Configuration) HealthcheckController {
	return &healthcheckController{
		cfg: cfg,
	}
}
