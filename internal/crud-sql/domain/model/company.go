package model

type Company struct {
	tableName struct{} `pg:"crud_sql.companies"`

	Id   *int64  `pg:"id,pk"`
	Name *string `pg:"name,notnull,unique"`
	Cnpj *string `pg:"cnpj,notnull,unique"`

	Jobs []*Job `pg:"rel:has-many"`

	Base
}
