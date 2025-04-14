package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/RakibulBh/fido2-microservice/internal/store"
	"github.com/go-redis/redis"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
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

	err = app.store.User.CreateUser(ctx, user)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	optionsWithSession := &OptionsWithSession{
		Options:    options,
		SessionKey: sessionKey,
	}

	fmt.Print(optionsWithSession)

	app.writeJSON(w, http.StatusOK, "registration successful", optionsWithSession)
}

type RegisterCompletePayload struct {
	ID         string   `json:"id"`
	RawID      string   `json:"rawId"`
	Response   response `json:"response"`
	Type       string   `json:"type"`
	SessionKey string   `json:"sessionKey"`
}

type response struct {
	AttestationObject string `json:"attestationObject"`
	ClientDataJSON    string `json:"clientDataJSON"`
}

func (app *application) CompleteRegister(w http.ResponseWriter, r *http.Request) {
	// Parse payload for validation
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("failed to read request body"))
		return
	}
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	// Parse payload for validation
	var payload RegisterCompletePayload
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		app.badRequestResponse(w, r, errors.New("invalid JSON payload"))
		return
	}

	ctx := r.Context()

	// Retrieve sessionData
	sessionBytes, err := app.redis.Get(ctx, payload.SessionKey).Bytes()
	if err == redis.Nil {
		app.badRequestResponse(w, r, errors.New("session expired or invalid"))
		return
	}
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	// Deserialize sessionData
	var sessionData webauthn.SessionData
	if err := json.Unmarshal(sessionBytes, &sessionData); err != nil {
		app.internalServerErrorResponse(w, r, errors.New("error deserializing session data"))
		return
	}

	// Retrieve user
	userID := string(sessionData.UserID)
	user, err := app.store.User.GetUserByID(ctx, userID)
	if err != nil {
		app.internalServerErrorResponse(w, r, errors.New("error getting user from db"))
		return
	}

	// Finish registration
	credential, err := app.webAuthnService.FinishRegistration(user, sessionData, r)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	// Store credential
	err = app.store.Auth.StoreCredential(ctx, user.ID, credential)
	if err != nil {
		app.errorJSON(w, errors.New("error storing credentials in database"), http.StatusInternalServerError)
		return
	}

	// Delete session
	app.redis.Del(ctx, payload.SessionKey)

	app.writeJSON(w, http.StatusCreated, "registration successful", nil)
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Login", nil)
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Logout", nil)
}
