package about

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(lang string, limit, offset int) ([]About, error)
	Count() (int64, error)
	FindByID(id uuid.UUID) (About, error)
	FindBySlug(slug string) (About, error)
	Create(p *About) error
	WithTx(fn func(tx *gorm.DB) error) error
	// Update(id uuid.UUID, p *Menu) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(lang string, limit, offset int) ([]About, error) {
	var abouts []About
	query := r.db.Model(&About{})

	if lang != "" {
		query = query.Preload("Translations", "lang = ?", lang)
	} else {
		query = query.Preload("Translations") // ✅ lấy full
	}
	if err := query.
		Limit(limit).Offset(offset).Find(&abouts).Error; err != nil {
		return nil, err
	}
	return abouts, nil
}

func (r *repository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&About{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) FindByID(id uuid.UUID) (About, error) {
	var about About
	if err := r.db.First(&about, "id = ?", id).Error; err != nil {
		return About{}, err
	}
	return about, nil
}

func (r *repository) Create(p *About) error {
	return r.db.Create(p).Error
}

func (r *repository) FindBySlug(slug string) (About, error) {
	var about About
	if err := r.db.First(&about, "slug = ?", slug).Error; err != nil {
		return About{}, err
	}
	return about, nil
}
func (r *repository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&About{}, "id = ?", id).Error
}
