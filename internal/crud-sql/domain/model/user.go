package model

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:crud_sql.users,alias:u"`

	Id    *int64  `bun:"id,pk,autoincrement"`
	Name  *string `bun:"name,"`
	Email *string `bun:"email,notnull,unique"`
	Rg    *string `bun:"rg,notnull,unique"`
	Cpf   *string `bun:"cpf,notnull,unique"`

	Jobs []*Job `bun:"rel:has-many"`

	Base
}

var _ bun.BeforeAppendModelHook = (*User)(nil)
