package store

import (
	"context"
	"database/sql"
)

type CredentialsStorage struct {
	db *sql.DB
}

type Credential struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	CredentialID string `json:"credential_id"`
	PublicKey    string `json:"public_key"`
	SignCount    string `json:"sign_count"`
}

func (s *CredentialsStorage) GetCredentialsByEmail(ctx context.Context, email string) (*Credential, error) {
	query := `
		SELECT id, email
		FROM users
		WHERE email = $1 
	`

	var fetchedUser User
	err := s.db.QueryRowContext(ctx, query, email).Scan(&fetchedUser.ID, &fetchedUser.Email)
	if err != nil {
		return nil, err
	}

	// fetch user credentials
	query = `
		SELECT id, user_id, credential_id, public_key, sign_count
		FROM webauthn_credentials
		WHERE user_id = $1
	`

	var userCredential Credential
	err = s.db.QueryRowContext(ctx, query, fetchedUser.ID).Scan(&userCredential.ID, &userCredential.UserID, &userCredential.CredentialID, &userCredential.PublicKey, &userCredential.SignCount)
	if err != nil {
		return nil, err
	}

	return &userCredential, nil
}
