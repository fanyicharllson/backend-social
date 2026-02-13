package main

import (
	"log"

	"github.com/fanyicharllson/backend-social/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config{
		addr: env.GetString("SERVER_ADDR", ":8081"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
