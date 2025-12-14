package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      string    `gorm:"default:user" json:"role"`
	FullName  string    `gorm:"not null" json:"name"`
	AvatarURL string    `gorm:"size:255" json:"avatar_url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
