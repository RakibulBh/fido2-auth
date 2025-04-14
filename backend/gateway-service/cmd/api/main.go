package main

import (
	"log"

	"github.com/RakibulBh/gateway-microservice/internal/env"
)

func main() {
	cfg := config{
		env:    env.GetString("ENV", "development"),
		addr:   ":" + env.GetString("PORT", "8081"),
		apiURL: env.GetString("API_URL", "http://localhost:8081"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
