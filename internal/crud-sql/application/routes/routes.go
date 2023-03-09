package routes

import (
	"github.com/gin-gonic/gin"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/controller"
)

func NewHealthcheckRoute(c *gin.Engine, ctrl controller.HealthcheckController) {
	c.GET("/healthcheck", ctrl.GetServerStatus)
}

func NewEnvRoute(c *gin.Engine, ctrl controller.EnvController) {
	c.GET("/env", ctrl.GetEnvByName)
}

func NewUserRoute(c *gin.Engine, ctrl controller.UserController) {
	c.GET("/users", ctrl.GetUsers)
	c.GET("/users/:id", ctrl.GetUserById)
}
