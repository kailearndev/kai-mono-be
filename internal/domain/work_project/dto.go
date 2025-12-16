package work_project

type WorkProjectTranslationDTO struct {
	Lang         string `json:"lang" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Slogan       string `json:"slogan" binding:"required"`
	Content      string `json:"content" binding:"required"`
	Subtitle     string `json:"subtitle" binding:"required"`
	ThumbnailAlt string `json:"thumbnail_alt"`
	ThumbnailUrl string `json:"thumbnail_url" binding:"required"`
	Duration     string `json:"duration" binding:"required"`
	ClientName   string `json:"client_name" binding:"required"`
	Services     string `json:"services" binding:"required"`
	Technologies string `json:"technologies" binding:"required"`
	ProjectUrl   string `json:"project_url" binding:"required"`
	ProjectName  string `json:"project_name" binding:"required"`
}

// GetLang implements validator.TransactionLike for language uniqueness checks.

type CreateWorkProjectDTO struct {
	Order        int                         `json:"order" binding:"required"`
	IsActive     bool                        `json:"is_active" binding:"required"`
	Thumbnail    string                      `json:"thumbnail" binding:"required"`
	Link         string                      `json:"link" binding:"required"`
	Slug         string                      `json:"slug" binding:"required"`
	Translations []WorkProjectTranslationDTO `json:"translations"`
}

type UpdateWorkProjectDTO struct {
	Order        *int                        `json:"order"`
	IsActive     *bool                       `json:"is_active"`
	Thumbnail    *string                     `json:"thumbnail"`
	Link         *string                     `json:"link"`
	Slug         *string                     `json:"slug"`
	Translations []WorkProjectTranslationDTO `json:"translations"`
}

func (t WorkProjectTranslationDTO) GetLang() string {
	return t.Lang
}
