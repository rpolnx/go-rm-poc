package model

type User struct {
	tableName struct{} `pg:"users"`

	Id    int64  `pg:"id,pk"`
	Name  string `pg:"name,"`
	Email string `pg:"email,"`

	BaseModel
}
