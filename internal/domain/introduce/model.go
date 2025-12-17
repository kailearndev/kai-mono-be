package introduce

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Introduce struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Order     int       `gorm:"not null;default:0" json:"order"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Avatar    string    `gorm:"size:255;not null" json:"avatar"`
	Facebook  string    `gorm:"size:255" json:"facebook"`
	Instagram string    `gorm:"size:255" json:"instagram"`
	Linkedin  string    `gorm:"size:255" json:"linkedin"`
	Github    string    `gorm:"size:255" json:"github"`

	Translations []IntroduceTranslation `gorm:"constraint:OnDelete:CASCADE" json:"translations"`
}

type IntroduceTranslation struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	IntroduceID uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_introduce_lang" json:"introduce_id"`
	Slogan      string         `gorm:"size:255;not null" json:"slogan"`
	Content     pq.StringArray `gorm:"type:text[]" json:"content"`
	Lang        string         `gorm:"size:10;not null;uniqueIndex:idx_introduce_lang" json:"lang"` // vi, en, ja
	Title       string         `gorm:"size:255;not null" json:"title"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`

	Introduce Introduce `gorm:"foreignKey:IntroduceID;references:ID" json:"-"`
}

func (m IntroduceTranslationDTO) GetLang() string {
	return m.Lang
}
