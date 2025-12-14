package product

type ProductRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	SKU         string   `json:"sku" validate:"required,alphanum"`
	Price       float64  `json:"price" validate:"required,gt=0"`
	Stock       int      `json:"stock" validate:"gte=0"`
	ImageURLs   []string `json:"image_urls" validate:"omitempty,dive,url"`
	Description string   `json:"description,omitempty"`
	CategoryID  *string  `json:"category_id,omitempty"` // optional
}
