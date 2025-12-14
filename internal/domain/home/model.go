package home

import (
	"time"

	"github.com/google/uuid"
)

type HomeHero struct {
	ID              uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	BackgroundImage string                `gorm:"type:text;not null" json:"background_image"`
	CreatedAt       time.Time             `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time             `gorm:"autoUpdateTime" json:"updated_at"`
	Translations    []HomeHeroTranslation `gorm:"foreignKey:HomeHeroID"`
}
type HomeHeroTranslation struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	HomeHeroID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_home_hero_locale"`
	Locale     string    `gorm:"type:varchar(5);not null;uniqueIndex:idx_home_hero_locale"`

	Title    string   `gorm:"type:text"`
	Subtitle string   `gorm:"type:text"`
	CtaText  string   `gorm:"type:text"`
	HomeHero HomeHero `gorm:"constraint:OnDelete:CASCADE;"`
}
