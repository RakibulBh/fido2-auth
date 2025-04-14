package main

import (
	"context"
	"log"

	"github.com/RakibulBh/fido2-microservice/internal/db"
	"github.com/RakibulBh/fido2-microservice/internal/env"
	"github.com/RakibulBh/fido2-microservice/internal/store"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/redis/go-redis/v9"
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
			RPID:          "localhost",
			RPOrigins:     []string{"http://localhost:5173"},
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

	// Init redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	// Store
	store := store.NewStorage(db)

	app := &application{
		config:          cfg,
		store:           store,
		webAuthnService: webAuthn,
		redis:           client,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
