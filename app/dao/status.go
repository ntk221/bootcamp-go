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

func (s *statusRepositoryImpl) DeleteStatusByID(ctx context.Context, statusID object.StatusID) error {
	query := `
		DELETE FROM status WHERE id = ?
	`

	_, err := s.db.ExecContext(ctx, query, statusID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *statusRepositoryImpl) GetStatusesByParams(ctx context.Context, maxID, sinceID, limit int) (*object.StatusCollection, error) {
	sc := object.NewStatusCollection([]object.Status{})

	query := `
        SELECT * FROM status
        WHERE id < ? AND id > ?
        LIMIT ?
    `

	rows, err := s.db.QueryxContext(ctx, query, maxID, sinceID, limit)
	if err != nil {
		log.Println(err)
		return &object.StatusCollection{}, err
	}
	for rows.Next() {
		status := object.Status{}
		err := rows.StructScan(&status)
		if err != nil {
			log.Println(err)
			return &object.StatusCollection{}, err
		}
		sc = sc.AddStatus(status)
	}

	return sc, nil
}
