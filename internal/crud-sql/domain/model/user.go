package model

import "time"

type User struct {
	tableName struct{} `pg:"crud_sql.users"`

	Id    *int64  `pg:"id,pk"`
	Name  *string `pg:"name,"`
	Email *string `pg:"email,notnull,unique"`
	Rg    *string `pg:"rg,notnull,unique"`
	Cpf   *string `pg:"cpf,notnull,unique"`

	Jobs []*Job `pg:"rel:has-many"`

	CreatedAt *time.Time `pg:"created_at,default:now()"`
	UpdatedAt *time.Time `pg:"updated_at,default:now()"`
	DeletedAt *time.Time `pg:"deleted_at"`
}
