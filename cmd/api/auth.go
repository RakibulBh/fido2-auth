package main

import (
	"net/http"
)

func (app *application) Register(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Register", nil)
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Login", nil)
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, "Logout", nil)
}
