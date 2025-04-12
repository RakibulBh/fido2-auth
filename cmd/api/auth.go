package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RakibulBh/studygroup-backend/internal/store"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/google/uuid"
)

type RegisterPayload struct {
	Email string `json:"email"`
}

type OptionsWithSession struct {
	Options    *protocol.CredentialCreation `json:"options"`
	SessionKey string                       `json:"sessionKey"`
}

func (app *application) BeginRegister(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload
	err := app.readJSON(r, &payload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	user := &store.User{
		ID:          uuid.NewString(),
		Email:       payload.Email,
		Credentials: nil,
	}

	options, sessionData, err := app.webAuthnService.BeginRegistration(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// store the session data in memory
	sessionKey := "webauthn:session:" + uuid.NewString()

	sessionBytes, err := json.Marshal(sessionData)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	err = app.redis.SetEx(ctx, sessionKey, sessionBytes, 5*time.Minute).Err()
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	err = app.store.Auth.CreateUser(ctx, user)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	optionsWithSession := &OptionsWithSession{
		Options:    options,
		SessionKey: sessionKey,
	}

	app.writeJSON(w, http.StatusOK, "registration successful", optionsWithSession)
}

type RegisterCompletePayload struct {
	ID       string   `json:"id"`
	RawID    string   `json:"rawId"`
	Response response `json:"response"`
	Type     string   `json:"type"`
}
type response struct {
	AttestationObject string `json:"attestationObject"`
	ClientDataJSON    string `json:"clientDataJSON"`
}

func (app *application) CompleteRegister(w http.ResponseWriter, r *http.Request) {
	var payload RegisterCompletePayload
	err := app.readJSON(r, payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Login", nil)
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Logout", nil)
}
