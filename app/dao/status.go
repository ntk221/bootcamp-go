package dao

import (
	"context"
	"github.com/jmoiron/sqlx"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

// NewStatus : Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (s status) PostStatus(ctx context.Context, status object.Status) (object.Status, error) {
	//TODO implement me
	panic("implement me")
}
