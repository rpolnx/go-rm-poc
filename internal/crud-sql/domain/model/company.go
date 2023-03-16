package model

import "github.com/uptrace/bun"

type Company struct {
	bun.BaseModel `bun:"table:crud_sql.companies,alias:c"`

	Id   *int64  `bun:"id,pk,autoincrement"`
	Name *string `bun:"name,notnull,unique"`
	Cnpj *string `bun:"cnpj,notnull,unique"`

	Jobs []*Job `bun:"rel:has-many"`

	Base
}

var _ bun.BeforeAppendModelHook = (*Company)(nil)
