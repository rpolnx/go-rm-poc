package model

import "time"

type Company struct {
	tableName struct{} `pg:"crud_sql.companies"`

	Id   *int64  `pg:"id,pk"`
	Name *string `pg:"name,notnull,unique"`
	Cnpj *string `pg:"cnpj,notnull,unique"`

	Jobs []*Job `pg:"rel:has-many"`

	CreatedAt *time.Time `pg:"created_at,default:now()"`
	UpdatedAt *time.Time `pg:"updated_at,default:now()"`
	DeletedAt *time.Time `pg:"deleted_at"`
}
