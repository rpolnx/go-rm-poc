package model

import "time"

type Job struct {
	tableName struct{} `pg:"crud_sql.jobs"`

	Id          *int64   `pg:"id,pk"`
	MonthSalary *float64 `pg:"month_salary,notnull"`
	HoursPerDay *int     `pg:"hours_per_day,notnull"`

	CompanyId *int64   `pg:"company_id"`
	Company   *Company `pg:"fk:company_id,rel:has-one"`

	UserId *int64 `pg:"user_id"`
	User   *User  `pg:"fk:user_id,rel:has-one"`

	CreatedAt *time.Time `pg:"created_at,default:now()"`
	UpdatedAt *time.Time `pg:"updated_at,default:now()"`
	DeletedAt *time.Time `pg:"deleted_at"`
}
