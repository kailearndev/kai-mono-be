package menu

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
	g := r.Group("/menus")
	{
		g.GET("/", h.ListMenus)
		g.GET("/:id", h.GetMenuByID)
		g.POST("/", h.CreateMenu)
		g.PATCH("/:id", h.UpdateMenu)
		g.DELETE("/:id", h.DeleteMenu)
	}
}

func (h *Handler) ListMenus(c *gin.Context) {
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

	items, total, err := h.service.ListMenus(lang, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": total,
	})

}

func (h *Handler) CreateMenu(c *gin.Context) {
	var menu CreateMenuDTO
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	p, err := h.service.CreateMenu(menu)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, p)
}

func (h *Handler) GetMenuByID(c *gin.Context) {
	id := c.Param("id")
	p, err := h.service.GetMenuByID(uuid.MustParse(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return

	}
	response.Success(c, p)
}
func (h *Handler) UpdateMenu(c *gin.Context) {
	id := c.Param("id")
	var menu UpdateMenuDTO
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	menuID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid menu id")
		return
	}
	updated, err := h.service.UpdateMenu(menuID, menu)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, updated)
}

func (h *Handler) DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteMenu(uuid.MustParse(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "menu deleted"})
}
