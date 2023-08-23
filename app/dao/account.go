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
	accountRepositoryImpl struct {
		db *sqlx.DB
	}
)

// NewAccount : Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &accountRepositoryImpl{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (ar *accountRepositoryImpl) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := ar.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (ar *accountRepositoryImpl) CreateUser(ctx context.Context, account *object.Account) error {
	_, err := ar.db.ExecContext(ctx, "insert into account (username, password_hash) values (?, ?)",
		account.Username,
		account.PasswordHash,
	)
	if err != nil {
		log.Printf("%v\n", err)
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (ar *accountRepositoryImpl) FindByID(ctx context.Context, userID object.AccountID) (*object.Account, error) {
	entity := new(object.Account)
	err := ar.db.QueryRowxContext(ctx, "select * from account where id = ?", userID).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
