package upload

import (
	"kai-mono-be/pkg/cloudstorage"
	"kai-mono-be/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// step 1: tạo handler upload image
type Handler struct {
	storage *cloudstorage.CloudFlyConfig
}

// NewHandler khởi tạo handler upload với cấu hình cloud storage

func NewHandler(s *cloudstorage.CloudFlyConfig) *Handler {
	return &Handler{storage: s}
}

// RegisterRoutes đăng ký các route liên quan đến upload
func (h *Handler) RegisterRoutes(rg *gin.Engine) {
	rg.POST("/upload", h.UploadImage)
}

func (h *Handler) UploadImage(c *gin.Context) {
	// Bước 1: Nhận file từ request
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Missing file")
		return
	}
	// Bước 2: Mở file để đọc
	src, err := file.Open()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Cannot open file")
		return
	}
	// Bước 3: Upload file lên CloudFly
	url, err := h.storage.UploadImage(src, file)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Upload failed: "+err.Error())
		return
	}
	response.Success(c, gin.H{"url": url})

}
