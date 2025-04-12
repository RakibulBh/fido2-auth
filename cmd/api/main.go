package main

import (
	"log"

	"github.com/RakibulBh/studygroup-backend/internal/db"
	"github.com/RakibulBh/studygroup-backend/internal/env"
	"github.com/RakibulBh/studygroup-backend/internal/store"
	"github.com/go-webauthn/webauthn/webauthn"
)

func main() {
	cfg := config{
		env:    env.GetString("ENV", "development"),
		addr:   ":" + env.GetString("PORT", "8080"),
		apiURL: env.GetString("API_URL", "http://localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/fido2?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 10),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 10),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "10s"),
		},
		webAuthn: &webauthn.Config{
			RPDisplayName: "FIDO2-Demo",
			RPID:          "github.com/RakibulBh/fido2-auth",
			RPOrigins:     []string{"github.com/RakibulBh/fido2-auth"},
		},
	}

	// WEBAUTHN
	webAuthn, err := webauthn.New(cfg.webAuthn)
	if err != nil {
		panic(err)
	}

	// Database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Store
	store := store.NewStorage(db)

	app := &application{
		config:          cfg,
		store:           store,
		webAuthnService: webAuthn,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
