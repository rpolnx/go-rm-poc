package model

import "github.com/go-pg/pg/v10"

type Job struct {
	tableName struct{} `pg:"crud_sql.jobs"`

	Id          *int64   `pg:"id,pk"`
	MonthSalary *float64 `pg:"month_salary,notnull"`
	HoursPerDay *int     `pg:"hours_per_day,notnull"`

	CompanyId *int64   `pg:"company2_id"`
	Company   *Company `pg:"fk:company2_id,rel:has-one"`

	UserId *int64 `pg:"user2_id"`
	User   *User  `pg:"fk:user2_id,rel:has-one"`

	Base
}

var (
	_ pg.BeforeInsertHook = (*Job)(nil)
	_ pg.BeforeUpdateHook = (*Job)(nil)
	_ pg.BeforeDeleteHook = (*Job)(nil)
)
