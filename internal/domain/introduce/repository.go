package introduce

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(lang string, limit, offset int) ([]Introduce, error)
	Count() (int64, error)
	FindByID(id uuid.UUID) (Introduce, error)
	FindBySlug(slug string) (Introduce, error)
	Create(p *Introduce) error
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

func (r *repository) FindAll(lang string, limit, offset int) ([]Introduce, error) {
	var introduces []Introduce
	query := r.db.Model(&Introduce{})

	if lang != "" {
		query = query.Preload("Translations", "lang = ?", lang)
	} else {
		query = query.Preload("Translations") // ✅ lấy full
	}
	if err := query.
		Limit(limit).Offset(offset).Find(&introduces).Error; err != nil {
		return nil, err
	}
	return introduces, nil
}

func (r *repository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&Introduce{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) FindByID(id uuid.UUID) (Introduce, error) {
	var introduce Introduce
	if err := r.db.First(&introduce, "id = ?", id).Error; err != nil {
		return Introduce{}, err
	}
	return introduce, nil
}

func (r *repository) Create(p *Introduce) error {
	return r.db.Create(p).Error
}

func (r *repository) FindBySlug(slug string) (Introduce, error) {
	var introduce Introduce
	if err := r.db.First(&introduce, "slug = ?", slug).Error; err != nil {
		return Introduce{}, err
	}
	return introduce, nil
}
func (r *repository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Introduce{}, "id = ?", id).Error
}
