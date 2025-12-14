package user

import (
	"kai-mono-be/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	// Register user-related routes here
	g := r.Group("/users")
	g.POST("/register", h.RegisterUser)
	g.POST("/login", h.LoginUser)
}

func (h *Handler) RegisterUser(c *gin.Context) {
	var req RegisterUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.service.RegisterUser(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "registration failed")
		return
	}

	response.Created(c, user)
}

func (h *Handler) LoginUser(c *gin.Context) {
	var req LoginUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.LoginUser(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}
	response.Success(c, gin.H{"token": token})
}
