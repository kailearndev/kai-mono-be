package main

import (
	"kai-mono-be/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("❌ Error loading .env file")
	}

	// load .env automatically in db.InitPostgres
	app.Run()
}
