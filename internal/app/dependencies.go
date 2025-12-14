package app

import (
	"log"

	"kai-mono-be/internal/domain/home"
	"kai-mono-be/internal/domain/product"
	"kai-mono-be/internal/domain/upload"
	"kai-mono-be/internal/domain/user"
	"kai-mono-be/pkg/cloudstorage"

	"gorm.io/gorm"
)

// Dependencies holds all application dependencies
type Dependencies struct {
	ProductHandler *product.Handler
	UserHandler    *user.Handler
	HomeHandler    *home.Handler
	UploadHandler  *upload.Handler
}

// InitDependencies sets up all repositories, services, and handlers
func InitDependencies(db *gorm.DB, config *Config) *Dependencies {
	// Initialize storage
	storage, err := cloudstorage.NewCloudFlyConfig(
		config.CloudFlyEndpoint,
		config.CloudFlyAccessKey,
		config.CloudFlySecretKey,
		config.CloudFlyBucket,
	)
	if err != nil {
		log.Fatalln("❌ failed to connect CloudFly:", err)
	}

	return &Dependencies{
		ProductHandler: initProductHandler(db),
		UserHandler:    initUserHandler(db),
		HomeHandler:    initHomeHandler(db),
		UploadHandler:  upload.NewHandler(storage),
	}
}

func initProductHandler(db *gorm.DB) *product.Handler {
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	return product.NewHandler(service)
}

func initUserHandler(db *gorm.DB) *user.Handler {
	repo := user.NewRepository(db)
	service := user.NewService(repo)
	return user.NewHandler(service)
}

func initHomeHandler(db *gorm.DB) *home.Handler {
	repo := home.NewRepository(db)
	service := home.NewService(repo)
	return home.NewHandler(service)
}
