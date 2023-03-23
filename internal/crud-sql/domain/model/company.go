package model

import "github.com/go-pg/pg/v10"

type Company struct {
	tableName struct{} `pg:"crud_sql.companies"`

	Id   *int64  `pg:"id,pk"`
	Name *string `pg:"name,notnull,unique"`
	Cnpj *string `pg:"cnpj,notnull,unique"`

	Jobs []*Job `pg:"rel:has-many"`

	Base
}

var (
	_ pg.BeforeInsertHook = (*Company)(nil)
	_ pg.BeforeUpdateHook = (*Company)(nil)
	_ pg.BeforeDeleteHook = (*Company)(nil)
)
