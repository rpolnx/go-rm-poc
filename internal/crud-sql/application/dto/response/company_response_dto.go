package response_dto

import (
	"github.com/jinzhu/copier"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	"time"
)

type CompanyResponseDTO struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name,omitempty"`
	Cnpj *string `json:"cnpj,omitempty"`

	Jobs []*JobResponseDTO `json:"jobs,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (dto *CompanyResponseDTO) FromEntity(company *model.Company) *CompanyResponseDTO {
	_ = copier.Copy(&dto, &company)

	return dto
}
