package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
)

type EnvController interface {
	GetEnvByName(c *gin.Context)
}

type envController struct {
	cfg *config.Configuration
}

func (ctrl *envController) GetEnvByName(c *gin.Context) {
	envName := c.Param("env_name")
	envValue := os.Getenv(envName)

	m := map[string]interface{}{"name": envName, "value": envValue}

	logrus.Info(ctrl.cfg.App.Profile)

	c.JSON(http.StatusOK, m)
}

func NewEnvController(cfg *config.Configuration) EnvController {
	return &envController{
		cfg: cfg,
	}
}
