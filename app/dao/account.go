package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// NewAccount : Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (r *account) CreateUser(ctx context.Context, account *object.Account) error {
	_, err := r.db.ExecContext(ctx, "insert into account (username, password_hash) values (?, ?)",
		account.Username,
		account.PasswordHash,
	)
	if err != nil {
		log.Printf("%v\n", err)
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (r *account) PostStatus(ctx context.Context, account *object.Account) error {
	statuses, err := account.GetStatuses()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	yetPosted := make([]*object.Status, 0)
	for _, status := range statuses {
		if !status.Posted {
			yetPosted = append(yetPosted, status)
			status.Posted = true
		}
	}

	if len(yetPosted) == 0 {
		return errors.New("no status")
	}

	for _, status := range yetPosted {
		_, err := r.db.ExecContext(ctx, "insert into status (account_id, content) values (?, ?)",
			status.AccountID,
			status.Content,
		)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}
