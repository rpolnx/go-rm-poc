package model

import "time"

type User struct {
	tableName struct{} `pg:"crud_sql.users"`

	Id    *int64  `pg:"id,pk"`
	Name  *string `pg:"name,"`
	Email *string `pg:"email,"`

	CreatedAt *time.Time `pg:"created_at,default:now()"`
	UpdatedAt *time.Time `pg:"updated_at,default:now()"`
	DeletedAt *time.Time `pg:"deleted_at"`
}
