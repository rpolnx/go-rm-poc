package routes

import (
	"github.com/gin-gonic/gin"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/controller"
)

func NewHealthcheckRoute(c *gin.Engine, ctrl controller.HealthcheckController) {
	c.GET("/healthcheck", ctrl.GetServerStatus)
}

func NewUserRoute(c *gin.Engine, ctrl controller.UserController) {
	c.GET("/users", ctrl.GetUsers)
	c.GET("/users/:id", ctrl.GetUserById)
	c.POST("/users", ctrl.CreateUser)
	c.PUT("/users/:id", ctrl.UpdateUser)
	c.DELETE("/users/:id", ctrl.DeleteUser)
}

func NewCompanyRoute(c *gin.Engine, ctrl controller.CompanyController) {
	c.GET("/companies", ctrl.GetCompanies)
	c.GET("/companies/:id", ctrl.GetCompanyById)
	c.POST("/companies", ctrl.CreateCompany)
	c.PUT("/companies/:id", ctrl.UpdateCompany)
	c.DELETE("/companies/:id", ctrl.DeleteCompany)
}
