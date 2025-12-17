package introduce

type IntroduceTranslationDTO struct {
	Lang    string   `json:"lang" binding:"required"`
	Title   string   `json:"title" binding:"required"`
	Slogan  string   `json:"slogan" binding:"required"`
	Content []string `json:"content" binding:"required"`
}

type CreateIntroduceDTO struct {
	Order        int                       `json:"order"`
	IsActive     bool                      `json:"is_active"`
	Avatar       string                    `json:"avatar"`
	Facebook     string                    `json:"facebook"`
	Instagram    string                    `json:"instagram"`
	Linkedin     string                    `json:"linkedin"`
	Github       string                    `json:"github"`
	Translations []IntroduceTranslationDTO `json:"translations"`
}

type UpdateIntroduceDTO struct {
	Order        *int                      `json:"order"`
	IsActive     *bool                     `json:"is_active"`
	Avatar       *string                   `json:"avatar"`
	Facebook     *string                   `json:"facebook"`
	Instagram    *string                   `json:"instagram"`
	Linkedin     *string                   `json:"linkedin"`
	Github       *string                   `json:"github"`
	Translations []IntroduceTranslationDTO `json:"translations"`
}
