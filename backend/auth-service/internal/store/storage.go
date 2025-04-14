package store

import (
	"context"
	"database/sql"

	"github.com/go-webauthn/webauthn/webauthn"
)

type Storage struct {
	User interface {
		CreateUser(ctx context.Context, user *User) error
		GetUserByID(ctx context.Context, ID string) (*User, error)
	}
	Auth interface {
		StoreCredential(ctx context.Context, userID string, credential *webauthn.Credential) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		User: &UserStorage{db: db},
		Auth: &AuthStorage{db: db},
	}
}
