package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/dto/request"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/dto/response"
	port "rpolnx.com.br/crud-sql/internal/crud-sql/domain/ports"
	"strconv"
)

type UserController interface {
	GetUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	cfg         *config.Configuration
	userService port.UserUseCase
}

func (c *userController) GetUsers(ginCtx *gin.Context) {
	users, err := c.userService.GetAllUsers()

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	response := make([]*response_dto.UserResponseDTO, 0)

	for _, user := range users {
		userResponseDTO := &response_dto.UserResponseDTO{}
		userResponseDTO = userResponseDTO.FromEntity(user)
		response = append(response, userResponseDTO)
	}

	ginCtx.JSON(http.StatusOK, response)
}

func (c *userController) GetUserById(ginCtx *gin.Context) {
	id, err := strconv.ParseInt(ginCtx.Param("id"), 10, 64)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	user, err := c.userService.GetOneUser(&id)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	userResponseDTO := &response_dto.UserResponseDTO{}
	userResponseDTO = userResponseDTO.FromEntity(user)

	ginCtx.JSON(http.StatusOK, userResponseDTO)
}

func (c *userController) CreateUser(ginCtx *gin.Context) {
	userDto := &request_dto.UserRequestDTO{}

	err := ginCtx.ShouldBindJSON(userDto)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	id, err := c.userService.CreateUser(userDto.ToEntity())

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	userResponseDTO := &response_dto.UserResponseDTO{
		Id: id,
	}

	ginCtx.JSON(http.StatusCreated, userResponseDTO)
}

func (c *userController) UpdateUser(ginCtx *gin.Context) {
	id, err := strconv.ParseInt(ginCtx.Param("id"), 10, 64)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	userDto := &request_dto.UserRequestDTO{}

	err = ginCtx.ShouldBindJSON(userDto)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	updatedId, err := c.userService.UpdateUser(&id, userDto.ToEntity())

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	userResponseDTO := &response_dto.UserResponseDTO{
		Id: updatedId,
	}

	ginCtx.JSON(http.StatusOK, userResponseDTO)
}

func (c *userController) DeleteUser(ginCtx *gin.Context) {
	id, err := strconv.ParseInt(ginCtx.Param("id"), 10, 64)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	err = c.userService.DeleteUser(&id)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	ginCtx.Status(http.StatusOK)
}

func NewUserController(cfg *config.Configuration, userService port.UserUseCase) UserController {
	return &userController{
		cfg:         cfg,
		userService: userService,
	}
}
