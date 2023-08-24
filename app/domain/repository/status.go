package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	PostStatus(ctx context.Context, status object.Status) (object.Status, error)
	FindStatusByID(ctx context.Context, statusID object.StatusID) (object.Status, error)
	DeleteStatusByID(ctx context.Context, statusID object.StatusID) error

	GetStatusesByParams(ctx context.Context, maxID, sinceID, limit int) (*object.StatusCollection, error)
}
