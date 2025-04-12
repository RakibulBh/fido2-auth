package store

import (
	"context"
	"database/sql"

	"github.com/go-webauthn/webauthn/webauthn"
)

type AuthStorage struct {
	db *sql.DB
}

type User struct {
	ID          string
	Email       string
	Credentials []webauthn.Credential
}

// WebAuthnID returns the user's unique ID as a byte slice.
func (u *User) WebAuthnID() []byte {
	return []byte(u.ID)
}

// WebAuthnName returns the user's login name (e.g., email).
func (u *User) WebAuthnName() string {
	return u.Email
}

// WebAuthnDisplayName returns the user's display name for UI prompts.
func (u *User) WebAuthnDisplayName() string {
	return u.Email // Or a nickname, if you have one
}

// WebAuthnCredentials returns the user's stored credentials.
func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (a *AuthStorage) CreateUser(ctx context.Context, user *User) error {
	query := `
		INSERT into users (id, email)
		VALUES ($1, $2)
	`

	_, err := a.db.ExecContext(ctx, query, user.ID, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthStorage) StoreSessionData(ctx context.Context, userID string, sessionData *webauthn.SessionData) error {
	// query := `
	// 	INSERT INTO webauthn_credentials (user_id, credential_id, public_key, sign_count)
	// 	VALUES ($1, $2, $3, $4)
	// `

	return nil
}

func (a *AuthStorage) StoreCredential(ctx context.Context, userID string, credential *webauthn.Credential) error {
	query := `
		INSERT INTO webauthn_credentials (user_id, credential_id, public_key, sign_count)
		VALUES ($1, $2, $3, $4)
	`

	_, err := a.db.Exec(query, userID, credential.ID, credential.PublicKey, credential.Authenticator.SignCount)

	return err
}
