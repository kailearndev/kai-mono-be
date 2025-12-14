package main

import "kai-mono-be/internal/app"

func main() {
	// load .env automatically in db.InitPostgres
	app.Run()
}
