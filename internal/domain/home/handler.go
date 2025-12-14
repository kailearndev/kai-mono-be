package home

import (
	"net/http"
	"strconv"

	"kai-mono-be/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	// Register product-related routes here
	g := r.Group("/home")
	{
		g.GET("/", h.ListHome)
		g.GET("/:id", h.GetHomeByID)
		g.POST("/", h.CreateHome)
		// g.PUT("/:id", h.UpdateHome)
		// g.DELETE("/:id", h.DeleteHome)
	}
}

func (h *Handler) ListHome(c *gin.Context) {
	// Implementation of listing Home
	lang := c.Query("lang")
	limit := 10
	offset := 0
	if l := c.Query("limit"); l != "" {
		// parse limit
		if v, err := strconv.Atoi(l); err == nil {
			limit = v
		}
	}
	if o := c.Query("offset"); o != "" {
		// parse offset
		if v, err := strconv.Atoi(o); err == nil {
			offset = v
		}
	}

	items, total, err := h.service.ListHome(lang, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": total,
	})

}

func (h *Handler) CreateHome(c *gin.Context) {
	var req CreateHomeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	created, err := h.service.CreateHome(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Created(c, created)
}

func (h *Handler) GetHomeByID(c *gin.Context) {
	id := c.Param("id")
	lang := c.GetHeader("Accept-Language")
	_ = lang // currently not used, but can be passed to service if needed
	p, err := h.service.GetHomeByID(lang, uuid.MustParse(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return

	}
	response.Success(c, p)
}

// func (h *Handler) UpdateHome(c *gin.Context) {
// 	id := c.Param("id")
// 	var home HeroRequest
// 	if err := c.ShouldBindJSON(&home); err != nil {
// 		response.Error(c, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	updated, err := h.service.UpdateHome(uuid.MustParse(id), home)
// 	if err != nil {
// 		response.Error(c, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	response.Success(c, updated)
// }

// func (h *Handler) DeleteHome(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := h.service.DeleteHome(uuid.MustParse(id)); err != nil {
// 		response.Error(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	response.Success(c, gin.H{"message": "home deleted"})
// }
