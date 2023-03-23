package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"golang.org/x/net/context"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	port "rpolnx.com.br/crud-sql/internal/crud-sql/domain/ports"
	"rpolnx.com.br/crud-sql/internal/crud-sql/infrastructure/repository/postgres"
)

type companyRepository struct {
	Db *bun.DB
}

func (svc *companyRepository) GetAllCompaniesOut() ([]*model.Company, error) {
	companies := make([]*model.Company, 0)

	err := svc.Db.NewSelect().
		Model((*model.Company)(nil)).
		Relation("Jobs").
		Scan(context.Background(), &companies)

	return companies, postgres.HandleDbError(err)
}

func (svc *companyRepository) GetOneCompanyOut(id *int64) (*model.Company, error) {
	company := &model.Company{
		Id: id,
	}

	err := svc.Db.NewSelect().
		Model(company).
		WherePK().
		Relation("Jobs").
		Scan(context.Background())

	return company, postgres.HandleDbError(err)
}

func (svc *companyRepository) CreateCompanyOut(company *model.Company) (*int64, error) {
	company.Id = nil

	insert, err := svc.Db.NewInsert().
		Model(company).
		Exec(context.Background())

	logrus.Debugf("Result from select %v", insert)

	if err != nil {
		return nil, postgres.HandleDbError(err)
	}

	rowsAffected, err := insert.RowsAffected()

	affected := rowsAffected
	logrus.Info(affected)

	return company.Id, err
}

func (svc *companyRepository) UpdateCompanyOut(id *int64, company *model.Company) (*int64, error) {
	company.Id = id

	updated, err := svc.Db.NewUpdate().
		Model(company).
		WherePK().
		OmitZero().
		Exec(context.Background())

	if err != nil {
		return nil, postgres.HandleDbError(err)
	}

	affected, err := updated.RowsAffected()

	logrus.Infof("Afftected %d", affected)

	if affected == 0 {
		return nil, &config.NotFoundError{Name: "Company"}
	}

	return company.Id, err
}

func (svc *companyRepository) DeleteCompanyOut(id *int64) error {
	company := &model.Company{Id: id}

	deleted, err := svc.Db.NewDelete().
		Model(company).
		WherePK().
		WherePK().
		Exec(context.Background())

	if err != nil {
		return err
	}

	affected, err := deleted.RowsAffected()
	logrus.Info(affected)

	if affected == 0 {
		return &config.NotFoundError{Name: "Company"}
	}

	return postgres.HandleDbError(err)
}

func NewCompanyRepository(db *bun.DB) port.CompanyPort {
	return &companyRepository{
		Db: db,
	}
}
