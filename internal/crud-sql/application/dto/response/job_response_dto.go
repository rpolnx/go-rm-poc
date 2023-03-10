package response_dto

import (
	"github.com/jinzhu/copier"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	"time"
)

type JobResponseDTO struct {
	Id          *int64   `json:"id"`
	MonthSalary *float64 `json:"month_salary,omitempty"`
	HoursPerDay *int     `json:"hours_per_day,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (dto *JobResponseDTO) FromEntity(user *model.User) *JobResponseDTO {
	_ = copier.Copy(&dto, &user)

	return dto
}
