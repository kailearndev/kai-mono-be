package menu

type MenuTranslationDTO struct {
	Lang  string `json:"lang" binding:"required"`
	Title string `json:"title" binding:"required"`
}

type CreateMenuDTO struct {
	Slug         string               `json:"slug" binding:"required"`
	Order        int                  `json:"order"`
	IsActive     bool                 `json:"is_active"`
	Translations []MenuTranslationDTO `json:"translations"`
}

type UpdateMenuDTO struct {
	Slug         *string              `json:"slug"`
	Order        *int                 `json:"order"`
	IsActive     *bool                `json:"is_active"`
	Translations []MenuTranslationDTO `json:"translations"`
}
