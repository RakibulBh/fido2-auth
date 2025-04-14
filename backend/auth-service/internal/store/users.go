package store

import (
	"context"
	"database/sql"

	"github.com/go-webauthn/webauthn/webauthn"
)

type UserStorage struct {
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

func (a *UserStorage) CreateUser(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (id, email)
		VALUES ($1, $2)
	`

	_, err := a.db.ExecContext(ctx, query, user.ID, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserStorage) GetUserByID(ctx context.Context, ID string) (*User, error) {
	query := `
		SELECT id, email
		FROM users
		WHERE id = $1
	`

	user := &User{}

	err := a.db.QueryRowContext(ctx, query, ID).Scan(&user.ID, &user.Email)
	if err != nil {
		return &User{}, err
	}

	return user, nil
}
