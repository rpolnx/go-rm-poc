package request_dto

import (
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	"time"

	"github.com/jinzhu/copier"
)

type UserRequestDTO struct {
	Id    *int64  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Rg    *string `json:"rg,omitempty"`
	Cpf   *string `json:"cpf,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (dto *UserRequestDTO) ToEntity() (user *model.User) {
	user = &model.User{}

	_ = copier.Copy(user, &dto)

	return user
}
