package object

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type (
	AccountID    = int64
	StatusID     = int64
	PasswordHash = string

	// Account account
	Account struct {
		// The internal ID of the account
		ID AccountID `json:"-"`

		// The username of the account
		Username string `json:"username,omitempty"`

		// The username of the account
		PasswordHash string `json:"-" db:"password_hash"`

		// The account's display name
		DisplayName *string `json:"display_name,omitempty" db:"display_name"`

		// URL to the avatar image
		Avatar *string `json:"avatar,omitempty"`

		// URL to the header image
		Header *string `json:"header,omitempty"`

		// Biography of user
		Note *string `json:"note,omitempty"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`

		Statuses []*Status `json:"statuses,omitempty"`
	}

	Status struct {
		ID        StatusID  `json:"id"`
		AccountID AccountID `json:"account_id"`
		Content   string    `json:"content"`
		Posted    bool
	}
)

// Check if given password is match to account's password
func (a *Account) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(pass)) == nil
}

// Hash password and set it to account object
func (a *Account) SetPassword(pass string) error {
	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("generate error: %w", err)
	}
	a.PasswordHash = passwordHash
	return nil
}

func generatePasswordHash(pass string) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", errors.WithStack(err))
	}
	return PasswordHash(hash), nil
}

func CreateStatus(accountID AccountID, content string) (*Status, error) {
	// TODO: validation?

	return &Status{
		AccountID: accountID,
		Content:   content,
		Posted:    false,
	}, nil
}

func (a *Account) SetStatus(content string) error {
	status, err := CreateStatus(a.ID, content)
	if err != nil {
		return err
	}
	a.Statuses = append(a.Statuses, status)
	return nil
}

func (a *Account) GetStatuses() ([]*Status, error) {
	if len(a.Statuses) == 0 {
		return nil, errors.New("no status")
	}
	return a.Statuses, nil
}
