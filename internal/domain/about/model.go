package about

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type About struct {
	ID           uuid.UUID          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt    time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	Avatar       string             `gorm:"size:255;not null" json:"avatar"`
	Order        int                `gorm:"not null;default:0" json:"order"`
	IsActive     bool               `gorm:"default:true" json:"is_active"`
	Translations []AboutTranslation `gorm:"constraint:OnDelete:CASCADE" json:"translations"`
}

type AboutTranslation struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	AboutID     uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_about_lang" json:"about_id"`
	Lang        string         `gorm:"size:10;not null;uniqueIndex:idx_about_lang" json:"lang"` // vi, en, ja
	Title       string         `gorm:"size:255;not null" json:"title"`
	Content     string         `gorm:"type:text" json:"content"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Educations  datatypes.JSON `gorm:"type:jsonb" json:"educations"`
	Experiences datatypes.JSON `gorm:"type:jsonb" json:"experiences"`
	Reviews     datatypes.JSON `gorm:"type:jsonb" json:"reviews"`
	About       About          `gorm:"foreignKey:AboutID;references:ID" json:"-"`
}

type Education struct {
	School   string `json:"school" binding:"required"`
	Degree   string `json:"degree" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

type Experience struct {
	Company  string `json:"company" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}
type Review struct {
	ReviewerName string `json:"reviewer_name" binding:"required"`
	Position     string `json:"position" binding:"required"`
	Content      string `json:"content" binding:"required"`
	Avatar       string `json:"avatar" binding:"required"`
}
