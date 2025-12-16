package work_project

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
	g := r.Group("/work-projects")
	{
		g.GET("/", h.ListWorkProjects)
		g.GET("/:id", h.GetWorkProjectByID)
		g.GET("/slug/:slug", h.GetWorkProjectBySlug)
		g.POST("/", h.CreateWorkProject)
		g.PATCH("/:id", h.UpdateWorkProject)
		g.DELETE("/:id", h.DeleteWorkProject)
	}
}

func (h *Handler) ListWorkProjects(c *gin.Context) {
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

	items, total, err := h.service.ListWorkProjects(lang, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": total,
	})

}

func (h *Handler) CreateWorkProject(c *gin.Context) {
	var workProject CreateWorkProjectDTO
	if err := c.ShouldBindJSON(&workProject); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	p, err := h.service.CreateWorkProject(workProject)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, p)
}

func (h *Handler) GetWorkProjectByID(c *gin.Context) {
	id := c.Param("id")
	p, err := h.service.GetWorkProjectByID(uuid.MustParse(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return

	}
	response.Success(c, p)
}
func (h *Handler) UpdateWorkProject(c *gin.Context) {
	id := c.Param("id")
	var workProject UpdateWorkProjectDTO
	if err := c.ShouldBindJSON(&workProject); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	workProjectID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid work project id")
		return
	}
	_, err = h.service.UpdateWorkProject(workProjectID, workProject)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, "Work project updated successfully")
}

func (h *Handler) DeleteWorkProject(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteWorkProject(uuid.MustParse(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "work project deleted"})
}

func (h *Handler) GetWorkProjectBySlug(c *gin.Context) {
	slug := c.Param("slug")
	p, err := h.service.GetWorkProjectBySlug(slug)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return

	}
	response.Success(c, p)
}
