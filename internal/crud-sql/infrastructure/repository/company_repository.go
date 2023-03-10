package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	port "rpolnx.com.br/crud-sql/internal/crud-sql/domain/ports"
	"rpolnx.com.br/crud-sql/internal/crud-sql/infrastructure/repository/postgres"
	"time"
)

type companyRepository struct {
	Db *pg.DB
}

func (svc *companyRepository) GetAllCompaniesOut() ([]*model.Company, error) {
	companies := make([]*model.Company, 0)

	err := svc.Db.Model(&companies).
		Where("deleted_at IS NULL").
		Select()

	return companies, postgres.HandleDbError(err)
}

func (svc *companyRepository) GetOneCompanyOut(id *int64) (*model.Company, error) {
	company := &model.Company{
		Id: id,
	}

	err := svc.Db.Model(company).
		WherePK().
		Select()

	return company, postgres.HandleDbError(err)
}

func (svc *companyRepository) CreateCompanyOut(company *model.Company) (*int64, error) {
	company.Id = nil

	insert, err := svc.Db.
		Model(company).
		Insert()

	if err != nil {
		return nil, postgres.HandleDbError(err)
	}

	affected := insert.RowsAffected()
	logrus.Info(affected)

	return company.Id, err
}

func (svc *companyRepository) UpdateCompanyOut(id *int64, company *model.Company) (*int64, error) {
	now := time.Now()
	company.Id = id
	company.UpdatedAt = &now

	updated, err := svc.Db.
		Model(company).
		WherePK().
		UpdateNotZero()

	if err != nil {
		return nil, postgres.HandleDbError(err)
	}

	affected := updated.RowsAffected()

	logrus.Infof("Afftected %d", affected)

	if affected == 0 {
		return nil, &config.NotFoundError{Name: "Company"}
	}

	return company.Id, err
}

func (svc *companyRepository) DeleteCompanyOut(id *int64) error {
	now := time.Now()
	company := &model.Company{
		Id:        id,
		UpdatedAt: &now,
		DeletedAt: &now,
	}

	deleted, err := svc.Db.
		Model(company).
		WherePK().
		UpdateNotZero()

	if err != nil {
		return err
	}

	affected := deleted.RowsAffected()
	logrus.Info(affected)

	if affected == 0 {
		return &config.NotFoundError{Name: "Company"}
	}

	return postgres.HandleDbError(err)
}

func NewCompanyRepository(db *pg.DB) port.CompanyPort {
	return &companyRepository{
		Db: db,
	}
}
