package model

import "github.com/go-pg/pg/v10"

type User struct {
	tableName struct{} `pg:"crud_sql.users"`

	Id    *int64  `pg:"id,pk"`
	Name  *string `pg:"name,"`
	Email *string `pg:"email,notnull,unique"`
	Rg    *string `pg:"rg,notnull,unique"`
	Cpf   *string `pg:"cpf,notnull,unique"`

	Jobs []*Job `pg:"rel:has-many"`

	Base
}

var (
	_ pg.BeforeInsertHook = (*User)(nil)
	_ pg.BeforeUpdateHook = (*User)(nil)
	_ pg.BeforeDeleteHook = (*User)(nil)
)
