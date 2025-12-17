package app

import (
	"log"

	"kai-mono-be/internal/domain/about"
	"kai-mono-be/internal/domain/introduce"
	"kai-mono-be/internal/domain/menu"
	"kai-mono-be/internal/domain/upload"
	"kai-mono-be/internal/domain/user"
	"kai-mono-be/internal/domain/work_project"
	"kai-mono-be/pkg/cloudstorage"

	"gorm.io/gorm"
)

// Dependencies holds all application dependencies
type Dependencies struct {
	UserHandler        *user.Handler
	UploadHandler      *upload.Handler
	MenuHandler        *menu.Handler
	IntroduceHandler   *introduce.Handler
	WorkProjectHandler *work_project.Handler
	AboutHandler       *about.Handler
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
		// ProductHandler: initProductHandler(db),
		UserHandler:        initUserHandler(db),
		UploadHandler:      upload.NewHandler(storage),
		MenuHandler:        initMenuHandler(db),
		IntroduceHandler:   initIntroduceHandler(db),
		WorkProjectHandler: initWorkProjectHandler(db),
		AboutHandler:       initAboutHandler(db),
	}
}

// func initProductHandler(db *gorm.DB) *product.Handler {
// 	repo := product.NewRepository(db)
// 	service := product.NewService(repo)
// 	return product.NewHandler(service)
// }

func initUserHandler(db *gorm.DB) *user.Handler {
	repo := user.NewRepository(db)
	service := user.NewService(repo)
	return user.NewHandler(service)
}

func initMenuHandler(db *gorm.DB) *menu.Handler {
	repo := menu.NewRepository(db)
	service := menu.NewService(repo)
	return menu.NewHandler(service)
}

func initIntroduceHandler(db *gorm.DB) *introduce.Handler {
	repo := introduce.NewRepository(db)
	service := introduce.NewService(repo)
	return introduce.NewHandler(service)
}

func initWorkProjectHandler(db *gorm.DB) *work_project.Handler {
	repo := work_project.NewRepository(db)
	service := work_project.NewService(repo)
	return work_project.NewHandler(service)
}

func initAboutHandler(db *gorm.DB) *about.Handler {
	repo := about.NewRepository(db)
	service := about.NewService(repo)
	return about.NewHandler(service)
}
