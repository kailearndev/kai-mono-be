package home

type CreateHomeRequest struct {
	BackgroundImage string                  `json:"background_image" binding:"required"`
	Translations    []CreateHomeTranslation `json:"translations" binding:"required,dive"`
}

type CreateHomeTranslation struct {
	Locale   string `json:"locale" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Subtitle string `json:"subtitle"`
	CtaText  string `json:"cta_text"`
}
