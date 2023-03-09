package model

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `pg:"default:now()"`
	UpdatedAt time.Time
	DeletedAt time.Time
}
