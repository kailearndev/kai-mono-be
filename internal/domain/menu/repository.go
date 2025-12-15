package menu

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(lang string, limit, offset int) ([]Menu, error)
	Count() (int64, error)
	FindByID(id uuid.UUID) (Menu, error)
	Create(p *Menu) error
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

func (r *repository) FindAll(lang string, limit, offset int) ([]Menu, error) {
	var menus []Menu
	query := r.db.Model(&Menu{})

	if lang != "" {
		query = query.Preload("Translations", "lang = ?", lang)
	} else {
		query = query.Preload("Translations") // ✅ lấy full
	}
	if err := query.
		Limit(limit).Offset(offset).Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
func (r *repository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&Menu{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (r *repository) FindByID(id uuid.UUID) (Menu, error) {
	var menu Menu
	if err := r.db.First(&menu, "id = ?", id).Error; err != nil {
		return Menu{}, err
	}
	return menu, nil
}

func (r *repository) Create(p *Menu) error {
	return r.db.Create(p).Error
}

func (r *repository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}
func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Menu{}, "id = ?", id).Error
}
