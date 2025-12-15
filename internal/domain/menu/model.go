package menu

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Slug     string    `gorm:"size:100;unique;not null" json:"slug"` // home, about, contact
	Order    int       `gorm:"not null;default:0" json:"order"`
	IsActive bool      `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Translations []MenuTranslation `gorm:"constraint:OnDelete:CASCADE" json:"translations"`
}

type MenuTranslation struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	MenuID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_menu_lang" json:"menu_id"`

	Lang      string    `gorm:"size:10;not null;uniqueIndex:idx_menu_lang" json:"lang"` // vi, en, ja
	Title     string    `gorm:"size:255;not null" json:"title"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Menu Menu `gorm:"foreignKey:MenuID;references:ID" json:"-"`
}
