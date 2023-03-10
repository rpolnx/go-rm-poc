package response_dto

import (
	"github.com/jinzhu/copier"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	"time"
)

type UserResponseDTO struct {
	Id    *int64  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (dto *UserResponseDTO) FromEntity(user *model.User) *UserResponseDTO {
	_ = copier.Copy(&dto, &user)

	return dto
}
