package service

import (
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	port "rpolnx.com.br/crud-sql/internal/crud-sql/domain/ports"
)

type companyService struct {
	companyPort port.CompanyPort
}

func (svc *companyService) GetAllCompanies() ([]*model.Company, error) {
	return svc.companyPort.GetAllCompaniesOut()
}

func (svc *companyService) GetOneCompany(id *int64) (*model.Company, error) {
	return svc.companyPort.GetOneCompanyOut(id)
}

func (svc *companyService) CreateCompany(user *model.Company) (*int64, error) {
	return svc.companyPort.CreateCompanyOut(user)
}

func (svc *companyService) UpdateCompany(id *int64, user *model.Company) (*int64, error) {
	return svc.companyPort.UpdateCompanyOut(id, user)
}

func (svc *companyService) DeleteCompany(id *int64) error {
	return svc.companyPort.DeleteCompanyOut(id)
}

func NewCompanyService(companyPort port.CompanyPort) port.CompanyUseCase {
	return &companyService{
		companyPort: companyPort,
	}
}
