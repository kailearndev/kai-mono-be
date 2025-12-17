package work_project

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type WorkProject struct {
	ID           uuid.UUID                `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt    time.Time                `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time                `gorm:"autoUpdateTime" json:"updated_at"`
	Thumbnail    string                   `gorm:"size:255;not null" json:"thumbnail"`
	Link         string                   `gorm:"size:255;not null" json:"link"`
	Slug         string                   `gorm:"size:255;not null;uniqueIndex" json:"slug"`
	Order        int                      `gorm:"not null;default:0" json:"order"`
	IsActive     bool                     `gorm:"default:true" json:"is_active"`
	Translations []WorkProjectTranslation `gorm:"constraint:OnDelete:CASCADE" json:"translations"`
}

type WorkProjectTranslation struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	WorkProjectID uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_work_project_lang" json:"work_project_id"`
	Slogan        string         `gorm:"size:255;not null" json:"slogan"`
	Content       string         `gorm:"type:text;not null" json:"content"`
	Lang          string         `gorm:"size:10;not null;uniqueIndex:idx_work_project_lang" json:"lang"` // vi, en, ja
	Title         string         `gorm:"size:255;not null" json:"title"`
	Subtitle      string         `gorm:"size:255;not null" json:"subtitle"`
	ThumbnailAlt  string         `gorm:"size:255;not null" json:"thumbnail_alt"`
	ThumbnailUrl  string         `gorm:"size:255;not null" json:"thumbnail_url"`
	Duration      string         `gorm:"size:100;not null" json:"duration"`
	ClientName    string         `gorm:"size:255;not null" json:"client_name"`
	Services      string         `gorm:"size:255;not null" json:"services"`
	Technologies  string         `gorm:"size:255;not null" json:"technologies"`
	ProjectUrl    string         `gorm:"size:255;not null" json:"project_url"`
	ProjectName   string         `gorm:"size:255;not null" json:"project_name"`
	Images        pq.StringArray `gorm:"type:text[]" json:"images"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`

	WorkProject WorkProject `gorm:"foreignKey:WorkProjectID;references:ID" json:"-"`
}
