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

type CompanyController interface {
	GetCompanies(c *gin.Context)
	GetCompanyById(c *gin.Context)
	CreateCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
}

type companyController struct {
	cfg            *config.Configuration
	companyService port.CompanyUseCase
}

func (c *companyController) GetCompanies(ginCtx *gin.Context) {
	companies, err := c.companyService.GetAllCompanies()

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	response := make([]*response_dto.CompanyResponseDTO, 0)

	for _, company := range companies {
		companyResponseDTO := &response_dto.CompanyResponseDTO{}
		companyResponseDTO = companyResponseDTO.FromEntity(company)
		response = append(response, companyResponseDTO)
	}

	ginCtx.JSON(http.StatusOK, response)
}

func (c *companyController) GetCompanyById(ginCtx *gin.Context) {
	id, err := strconv.ParseInt(ginCtx.Param("id"), 10, 64)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	company, err := c.companyService.GetOneCompany(&id)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	companyResponseDTO := &response_dto.CompanyResponseDTO{}
	companyResponseDTO = companyResponseDTO.FromEntity(company)

	ginCtx.JSON(http.StatusOK, companyResponseDTO)
}

func (c *companyController) CreateCompany(ginCtx *gin.Context) {
	companyDto := &request_dto.CompanyRequestDTO{}

	err := ginCtx.ShouldBindJSON(companyDto)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	id, err := c.companyService.CreateCompany(companyDto.ToEntity())

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	companyResponseDto := &response_dto.CompanyResponseDTO{
		Id: id,
	}

	ginCtx.JSON(http.StatusCreated, companyResponseDto)
}

func (c *companyController) UpdateCompany(ginCtx *gin.Context) {
	id, err := strconv.ParseInt(ginCtx.Param("id"), 10, 64)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	companyDto := &request_dto.CompanyRequestDTO{}

	err = ginCtx.ShouldBindJSON(companyDto)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	updatedId, err := c.companyService.UpdateCompany(&id, companyDto.ToEntity())

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	responseDTO := &response_dto.UserResponseDTO{
		Id: updatedId,
	}

	ginCtx.JSON(http.StatusOK, responseDTO)
}

func (c *companyController) DeleteCompany(ginCtx *gin.Context) {
	id, err := strconv.ParseInt(ginCtx.Param("id"), 10, 64)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err, http.StatusBadRequest))
		return
	}

	err = c.companyService.DeleteCompany(&id)

	if err != nil {
		ginCtx.JSON(config.HandleHttpError(ginCtx, err))
		return
	}

	ginCtx.Status(http.StatusOK)
}

func NewCompanyController(cfg *config.Configuration, companyService port.CompanyUseCase) CompanyController {
	return &companyController{
		cfg:            cfg,
		companyService: companyService,
	}
}
