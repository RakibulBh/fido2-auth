package store

import (
	"context"
	"database/sql"

	"github.com/go-webauthn/webauthn/webauthn"
)

type AuthStorage struct {
	db *sql.DB
}

func (a *AuthStorage) StoreCredential(ctx context.Context, userID string, credential *webauthn.Credential) error {
	query := `
		INSERT INTO webauthn_credentials (user_id, credential_id, public_key, sign_count)
		VALUES ($1, $2, $3, $4)
	`

	_, err := a.db.Exec(query, userID, credential.ID, credential.PublicKey, credential.Authenticator.SignCount)

	return err
}
