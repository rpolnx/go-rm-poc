package port

import "rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"

type CompanyUseCase interface {
	GetAllCompanies() ([]*model.Company, error)
	GetOneCompany(id *int64) (*model.Company, error)
	CreateCompany(user *model.Company) (*int64, error)
	UpdateCompany(id *int64, user *model.Company) (*int64, error)
	DeleteCompany(id *int64) error
}

type CompanyPort interface {
	GetAllCompaniesOut() ([]*model.Company, error)
	GetOneCompanyOut(ID *int64) (*model.Company, error)
	CreateCompanyOut(user *model.Company) (*int64, error)
	UpdateCompanyOut(id *int64, user *model.Company) (*int64, error)
	DeleteCompanyOut(id *int64) error
}
