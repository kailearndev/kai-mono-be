package about

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
	g := r.Group("/about")
	{
		g.GET("/", h.ListAbouts)
		g.GET("/:id", h.GetAboutByID)
		g.POST("/", h.CreateAbout)
		g.PATCH("/:id", h.UpdateAbout)
		g.DELETE("/:id", h.DeleteAbout)
	}
}

func (h *Handler) ListAbouts(c *gin.Context) {
	// Implementation of listing products
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

	items, total, err := h.service.ListAbouts(lang, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": total,
	})

}

func (h *Handler) CreateAbout(c *gin.Context) {
	var about CreateAboutDTO
	if err := c.ShouldBindJSON(&about); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	p, err := h.service.CreateAbout(about)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, p)
}

func (h *Handler) GetAboutByID(c *gin.Context) {
	id := c.Param("id")

	p, err := h.service.GetAboutByID(uuid.MustParse(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return

	}
	response.Success(c, p)
}
func (h *Handler) UpdateAbout(c *gin.Context) {
	id := c.Param("id")
	var about UpdateAboutDTO
	if err := c.ShouldBindJSON(&about); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	aboutID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid about id")
		return
	}
	_, err = h.service.UpdateAbout(aboutID, about)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, "Work project updated successfully")
}

func (h *Handler) DeleteAbout(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteAbout(uuid.MustParse(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "about deleted"})
}
