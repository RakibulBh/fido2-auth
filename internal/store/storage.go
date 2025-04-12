package store

import (
	"context"
	"database/sql"

	"github.com/go-webauthn/webauthn/webauthn"
)

type Storage struct {
	Auth interface {
		CreateUser(ctx context.Context, user *User) error
		StoreSessionData(ctx context.Context, userID string, sessionData *webauthn.SessionData) error
		StoreCredential(ctx context.Context, userID string, credential *webauthn.Credential) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Auth: &AuthStorage{db: db},
	}
}
