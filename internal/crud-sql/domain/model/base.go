package model

import (
	"context"
	"time"
)

type Base struct {
	tableName struct{} `pg:"-"`

	CreatedAt *time.Time `pg:"created_at,default:now()"`
	UpdatedAt *time.Time `pg:"updated_at,default:now()"`
	DeletedAt *time.Time `pg:"deleted_at,soft_delete"`
}

func (u *Base) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now()
	u.CreatedAt = &(now)

	return ctx, nil
}

func (u *Base) BeforeUpdate(ctx context.Context) (context.Context, error) {
	now := time.Now()

	u.UpdatedAt = &(now)

	return ctx, nil
}

func (u *Base) BeforeDelete(ctx context.Context) (context.Context, error) {
	now := time.Now()

	u.DeletedAt = &(now)

	return ctx, nil
}
