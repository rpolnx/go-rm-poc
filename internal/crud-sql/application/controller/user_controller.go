package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	port "rpolnx.com.br/crud-sql/internal/crud-sql/domain/ports"
	"strconv"
)

type UserController interface {
	GetUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userController struct {
	cfg         *config.Configuration
	userService port.UserUseCase
}

func (c *userController) GetUsers(ginCtx *gin.Context) {
	res, err := c.userService.GetAllUsers()

	logrus.Error(err)

	ginCtx.JSON(http.StatusOK, res)
}

func (c *userController) GetUserById(ginCtx *gin.Context) {
	id, err := strconv.ParseInt(ginCtx.Param("id"), 10, 64)

	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, config.HandleError(ginCtx, http.StatusBadRequest, err))
	}

	res, err := c.userService.GetOneUser(id)

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, config.HandleError(ginCtx, http.StatusInternalServerError, err))
	}

	ginCtx.JSON(http.StatusOK, res)
}

func (c *userController) CreateUser(ginCtx *gin.Context) {

	user := &model.User{}

	err := ginCtx.ShouldBindJSON(user)

	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, config.HandleError(ginCtx, http.StatusBadRequest, err))
	}

	res, err := c.userService.CreateUser(user)

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, config.HandleError(ginCtx, http.StatusInternalServerError, err))
	}

	ginCtx.JSON(http.StatusCreated, map[string]interface{}{
		"Id": res,
	})
}

func NewUserController(cfg *config.Configuration, userService port.UserUseCase) UserController {
	return &userController{
		cfg:         cfg,
		userService: userService,
	}
}
