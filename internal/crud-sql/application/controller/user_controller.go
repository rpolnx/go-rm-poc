package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
)

type UserController interface {
	GetUsers(c *gin.Context)
	GetUserById(c *gin.Context)
}

type userController struct {
	cfg *config.Configuration
}

func (h *userController) GetUsers(c *gin.Context) {
	m := gin.H{"status": "OK"}

	c.JSON(http.StatusOK, m)
}

func (h *userController) GetUserById(c *gin.Context) {
	m := map[string]interface{}{"status": "OK"}

	c.JSON(http.StatusOK, m)
}

func NewUserController(cfg *config.Configuration) UserController {
	return &userController{
		cfg: cfg,
	}
}
