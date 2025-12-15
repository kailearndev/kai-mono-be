package app

import (
	"log"

	"kai-mono-be/internal/domain/menu"
	"kai-mono-be/internal/domain/product"
	"kai-mono-be/internal/domain/user"
	"kai-mono-be/pkg/db"

	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	database := db.InitPostgres()

	// Enable UUID extension
	if err := database.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-generate-v4()"`).Error; err != nil {
		log.Println("⚠️  pgcrypto extension already exists or failed to create:", err)
	}

	// Run migrations
	if err := autoMigrate(database); err != nil {
		log.Fatalln("❌ migrate failed:", err)
	}

	log.Println("✅ Database initialized and migrated")
	return database
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&product.Product{},
		&user.User{},
		&menu.Menu{},
		&menu.MenuTranslation{},
	)
}
