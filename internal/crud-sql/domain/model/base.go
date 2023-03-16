package model

import (
	"github.com/uptrace/bun"
	"golang.org/x/net/context"
	"time"
)

type Base struct {
	bun.BaseModel `bun:"-"`

	CreatedAt *time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete"`
}

func (u *Base) BeforeAppendModel(_ context.Context, query bun.Query) error {
	now := time.Now()

	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = &(now)
	case *bun.UpdateQuery:
		u.UpdatedAt = &(now)
	case *bun.DeleteQuery:
		u.DeletedAt = &(now)
	}
	return nil
}
