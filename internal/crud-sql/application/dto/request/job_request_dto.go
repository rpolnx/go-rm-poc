package request_dto

import (
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	"time"

	"github.com/jinzhu/copier"
)

type JobRequestDTO struct {
	Id          *int64   `json:"id"`
	MonthSalary *float64 `json:"month_salary,omitempty"`
	HoursPerDay *int     `json:"hours_per_day,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (dto *JobRequestDTO) ToEntity() (company *model.Job) {
	company = &model.Job{}

	_ = copier.Copy(company, &dto)

	return company
}
