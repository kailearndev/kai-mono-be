package product

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
	g := r.Group("/products")
	{
		g.GET("/", h.ListProducts)
		g.GET("/:id", h.GetProductByID)
		g.POST("/", h.CreateProduct)
		g.PUT("/:id", h.UpdateProduct)
		g.DELETE("/:id", h.DeleteProduct)
	}
}

func (h *Handler) ListProducts(c *gin.Context) {
	// Implementation of listing products
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

	items, total, err := h.service.ListProducts(limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": total,
	})

}

func (h *Handler) CreateProduct(c *gin.Context) {
	var product ProductRequest
	if err := c.ShouldBindJSON(&product); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	p, err := h.service.CreateProduct(product)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, p)
}

func (h *Handler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	p, err := h.service.GetProductByID(uuid.MustParse(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return

	}
	response.Success(c, p)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product ProductRequest
	if err := c.ShouldBindJSON(&product); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	updated, err := h.service.UpdateProduct(uuid.MustParse(id), product)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, updated)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteProduct(uuid.MustParse(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "product deleted"})
}
