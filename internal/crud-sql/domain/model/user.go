package model

import "time"

type User struct {
	tableName struct{} `pg:"crud_sql.users"`

	Id    *int64  `json:"id" pg:"id,pk"`
	Name  *string `json:"name,omitempty" pg:"name,"`
	Email *string `json:"email,omitempty" pg:"email,"`

	CreatedAt *time.Time `json:"created_at,omitempty" pg:"created_at,default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" pg:"updated_at,default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" pg:"deleted_at"`
}
