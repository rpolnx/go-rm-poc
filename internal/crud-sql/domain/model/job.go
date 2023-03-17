package model

import "github.com/uptrace/bun"

type Job struct {
	bun.BaseModel `bun:"table:crud_sql.jobs,alias:j"`

	Id          *int64   `bun:"id,pk,autoincrement"`
	MonthSalary *float64 `bun:"month_salary,notnull"`
	HoursPerDay *int     `bun:"hours_per_day,notnull"`

	CompanyId *int64   `bun:"company_id"`
	Company   *Company `bun:"rel:has-one,join:company_id=id"`

	UserId *int64 `bun:"user_id"`
	User   *User  `bun:"rel:belongs-to,join:user_id=id"`

	Base
}

var _ bun.BeforeAppendModelHook = (*Job)(nil)
