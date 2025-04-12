package main

import (
	"net/http"

	"github.com/RakibulBh/studygroup-backend/internal/store"
	"github.com/google/uuid"
)

type RegisterPayload struct {
	Email string `json:"email"`
}

func (app *application) Register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload
	err := app.readJSON(r, &payload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	user := &store.User{
		ID:    uuid.NewString(),
		Email: payload.Email,
	}

	options, _, err := app.webAuthnService.BeginRegistration(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// TODO: Store session data
	// app.store.Auth.StoreSessionData(ctx, sessionData)

	err = app.store.Auth.CreateUser(ctx, user)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusOK, "registration successful", options)
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Login", nil)
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Logout", nil)
}
