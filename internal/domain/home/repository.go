package home

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(lang string, limit, offset int) ([]HomeHero, error)
	Count() (int64, error)
	FindByID(lang string, id uuid.UUID) (HomeHero, error)
	Create(p *HomeHero) error
	// Update(p *HomeHero) error
	// Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(lang string, limit, offset int) ([]HomeHero, error) {
	var homeHeroes []HomeHero
	if err := r.db.Preload("Translations", "locale = ?", lang).
		Limit(limit).Offset(offset).Find(&homeHeroes).Error; err != nil {
		return nil, err
	}
	return homeHeroes, nil
}
func (r *repository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&HomeHero{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (r *repository) FindByID(lang string, id uuid.UUID) (HomeHero, error) {
	var homeHero HomeHero
	if err := r.db.Preload("Translations", "locale = ?", lang).First(&homeHero, "id = ?", id).Error; err != nil {
		return HomeHero{}, err
	}
	return homeHero, nil
}

func (r *repository) Create(hero *HomeHero) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		// 1️⃣ Xóa hero cũ (cascade translations)
		if err := tx.Exec("DELETE FROM home_heros").Error; err != nil {
			return err
		}

		// 2️⃣ Insert hero mới + translations
		if err := tx.Create(hero).Error; err != nil {
			return err
		}

		return nil
	})
}

// func (r *repository) Update(p *HomeHero) error {
// 	return r.db.Save(p).Error
// }
// func (r *repository) Delete(id uuid.UUID) error {
// 	return r.db.Delete(&HomeHero{}, "id = ?", id).Error
// }
