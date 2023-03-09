package controller

import (
	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, m)
}

func NewHealthcheckController(cfg *config.Configuration) HealthcheckController {
	return &healthcheckController{
		cfg: cfg,
	}
}
