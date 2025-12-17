package app

import (
	"log"

	"kai-mono-be/internal/server"
)

// Run is the main entry point for the backend app
func Run() {
	// 1. Load configuration
	config := LoadConfig()
	config.Validate()

	// 2. Initialize database
	database := initDatabase()

	// 3. Initialize all dependencies (repos, services, handlers)
	deps := InitDependencies(database, config)

	// 4. Setup router with all handlers
	router := server.SetupRouter(server.RouteConfig{
		UploadHandler:      deps.UploadHandler,
		UserHandler:        deps.UserHandler,
		MenuHandler:        deps.MenuHandler,
		IntroduceHandler:   deps.IntroduceHandler,
		WorkProjectHandler: deps.WorkProjectHandler,
		AboutHandler:       deps.AboutHandler,
	})

	// 5. Start server
	log.Printf("🚀 Server running at http://localhost:%s", config.Port)
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatalln("❌ failed to start server:", err)
	}
}
