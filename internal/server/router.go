package server

import (
	"kai-mono-be/internal/domain/about"
	"kai-mono-be/internal/domain/introduce"
	"kai-mono-be/internal/domain/menu"
	"kai-mono-be/internal/domain/product"
	"kai-mono-be/internal/domain/upload"
	"kai-mono-be/internal/domain/user"
	"kai-mono-be/internal/domain/work_project"

	"github.com/gin-gonic/gin"
)

// RouteConfig gom các handler lại để dễ truyền vào SetupRouter
type RouteConfig struct {
	ProductHandler     *product.Handler
	UploadHandler      *upload.Handler
	UserHandler        *user.Handler
	MenuHandler        *menu.Handler
	IntroduceHandler   *introduce.Handler
	WorkProjectHandler *work_project.Handler
	AboutHandler       *about.Handler

	// sau này có thể thêm:
	// CategoryHandler *category.Handler
	// AuthHandler     *auth.Handler
}

// SetupRouter initializes all routes
func SetupRouter(cfg RouteConfig) *gin.Engine {
	r := gin.Default()

	// Register domain routes
	if cfg.ProductHandler != nil {
		cfg.ProductHandler.RegisterRoutes(r)
	}

	if cfg.UploadHandler != nil {
		cfg.UploadHandler.RegisterRoutes(r)
	}

	if cfg.UserHandler != nil {
		cfg.UserHandler.RegisterRoutes(r)
	}

	if cfg.MenuHandler != nil {
		cfg.MenuHandler.RegisterRoutes(r)
	}
	if cfg.IntroduceHandler != nil {
		cfg.IntroduceHandler.RegisterRoutes(r)
	}
	if cfg.WorkProjectHandler != nil {
		cfg.WorkProjectHandler.RegisterRoutes(r)
	}
	if cfg.AboutHandler != nil {
		cfg.AboutHandler.RegisterRoutes(r)
	}
	// future:
	// if cfg.CategoryHandler != nil {
	//     cfg.CategoryHandler.RegisterRoutes(r)
	// }

	return r
}
