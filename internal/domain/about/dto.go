package about

type AboutTranslationDTO struct {
	Lang        string       `json:"lang" binding:"required"`
	Title       string       `json:"title" binding:"required"`
	Content     string       `json:"content" binding:"required"`
	Educations  []Education  `json:"educations" binding:"required"`
	Experiences []Experience `json:"experiences" binding:"required"`
	Reviews     []Review     `json:"reviews" binding:"required"`
}

// GetLang implements validator.TransactionLike for language uniqueness checks.

type CreateAboutDTO struct {
	Order        int                   `json:"order" binding:"required"`
	IsActive     bool                  `json:"is_active" binding:"required"`
	Avatar       string                `json:"avatar" binding:"required"`
	Translations []AboutTranslationDTO `json:"translations"`
}

type UpdateAboutDTO struct {
	Avatar       *string                     `json:"avatar"`
	Order        *int                        `json:"order"`
	IsActive     *bool                       `json:"is_active"`
	Translations []UpdateAboutTranslationDTO `json:"translations"`
}

type UpdateAboutTranslationDTO struct {
	Lang        string       `json:"lang"`
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	Educations  []Education  `json:"educations"`
	Experiences []Experience `json:"experiences"`
	Reviews     []Review     `json:"reviews"`
}

func (t AboutTranslationDTO) GetLang() string {
	return t.Lang
}

func (t UpdateAboutTranslationDTO) GetLang() string {
	return t.Lang
}
