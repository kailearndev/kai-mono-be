package introduce

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
	g := r.Group("/introduces")
	{
		g.GET("/", h.ListIntroduces)
		g.GET("/:id", h.GetIntroduceByID)
		g.POST("/", h.CreateIntroduce)
		g.PATCH("/:id", h.UpdateIntroduce)
		g.DELETE("/:id", h.DeleteIntroduce)
	}
}

func (h *Handler) ListIntroduces(c *gin.Context) {
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

	items, total, err := h.service.ListIntroduces(lang, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": total,
	})

}

func (h *Handler) CreateIntroduce(c *gin.Context) {
	var introduce CreateIntroduceDTO
	if err := c.ShouldBindJSON(&introduce); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	p, err := h.service.CreateIntroduce(introduce)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, p)
}

func (h *Handler) GetIntroduceByID(c *gin.Context) {
	id := c.Param("id")
	p, err := h.service.GetIntroduceByID(uuid.MustParse(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return

	}
	response.Success(c, p)
}
func (h *Handler) UpdateIntroduce(c *gin.Context) {
	id := c.Param("id")
	var introduce UpdateIntroduceDTO
	if err := c.ShouldBindJSON(&introduce); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	introduceID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid introduce id")
		return
	}
	_, err = h.service.UpdateIntroduce(introduceID, introduce)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, "Introduce updated successfully")
}

func (h *Handler) DeleteIntroduce(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteIntroduce(uuid.MustParse(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "introduce deleted"})
}
