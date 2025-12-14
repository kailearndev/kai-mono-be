package server

import (
	"kai-mono-be/internal/domain/home"
	"kai-mono-be/internal/domain/product"
	"kai-mono-be/internal/domain/upload"
	"kai-mono-be/internal/domain/user"

	"github.com/gin-gonic/gin"
)

// RouteConfig gom các handler lại để dễ truyền vào SetupRouter
type RouteConfig struct {
	ProductHandler *product.Handler
	UploadHandler  *upload.Handler
	UserHandler    *user.Handler
	HomeHandler    *home.Handler
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
	if cfg.HomeHandler != nil {
		cfg.HomeHandler.RegisterRoutes(r)
	}
	// future:
	// if cfg.CategoryHandler != nil {
	//     cfg.CategoryHandler.RegisterRoutes(r)
	// }

	return r
}
