package dao

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"
)

type (
	// Implementation for repository.Account
	statusRepositoryImpl struct {
		db *sqlx.DB
	}
)

// NewStatus : Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &statusRepositoryImpl{db: db}
}

func (s *statusRepositoryImpl) PostStatus(ctx context.Context, status object.Status) (object.Status, error) {
	now := object.DateTime{Time: time.Now()}
	query := `
		insert into status (account_id, content, create_at)
		values (?, ?, ?);
	`

	accountID := status.AccountID
	content := status.Content

	result, err := s.db.ExecContext(ctx,
		query,
		accountID,
		content,
		now.Time,
	)
	if err != nil {
		return object.Status{}, err
	}

	insertedID, _ := result.LastInsertId()
	insertedStatus := object.Status{
		ID:        insertedID,
		AccountID: accountID,
		Content:   content,
		CreateAt:  now,
	}
	return insertedStatus, nil
}

//  `id` bigint(20) NOT NULL AUTO_INCREMENT,
//  `account_id` bigint(20) NOT NULL,
//  `content` text NOT NULL,
//  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

func (s *statusRepositoryImpl) FindStatusByID(ctx context.Context, statusID object.StatusID) (object.Status, error) {
	// TODO: implement me
	var entity object.Status

	query := `
		SELECT * FROM status WHERE id = ?
	`

	err := s.db.QueryRowxContext(ctx, query, statusID).StructScan(&entity)
	if err != nil {
		log.Println(err)
		return object.Status{}, err
	}

	log.Println(entity)

	return entity, nil
}
