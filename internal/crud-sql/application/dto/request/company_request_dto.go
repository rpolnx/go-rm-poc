package request_dto

import (
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	"time"

	"github.com/jinzhu/copier"
)

type CompanyRequestDTO struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name,omitempty"`
	Cnpj *string `json:"cnpj,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (dto *CompanyRequestDTO) ToEntity() (company *model.Company) {
	company = &model.Company{}

	_ = copier.Copy(company, &dto)

	return company
}
