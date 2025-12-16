package work_project

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(lang string, limit, offset int) ([]WorkProject, error)
	Count() (int64, error)
	FindByID(id uuid.UUID) (WorkProject, error)
	FindBySlug(slug string) (WorkProject, error)
	Create(p *WorkProject) error
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

func (r *repository) FindAll(lang string, limit, offset int) ([]WorkProject, error) {
	var workProjects []WorkProject
	query := r.db.Model(&WorkProject{})

	if lang != "" {
		query = query.Preload("Translations", "lang = ?", lang)
	} else {
		query = query.Preload("Translations") // ✅ lấy full
	}
	if err := query.
		Limit(limit).Offset(offset).Find(&workProjects).Error; err != nil {
		return nil, err
	}
	return workProjects, nil
}

func (r *repository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&WorkProject{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) FindByID(id uuid.UUID) (WorkProject, error) {
	var workProject WorkProject
	if err := r.db.First(&workProject, "id = ?", id).Error; err != nil {
		return WorkProject{}, err
	}
	return workProject, nil
}

func (r *repository) Create(p *WorkProject) error {
	return r.db.Create(p).Error
}

func (r *repository) FindBySlug(slug string) (WorkProject, error) {
	var workProject WorkProject
	if err := r.db.First(&workProject, "slug = ?", slug).Error; err != nil {
		return WorkProject{}, err
	}
	return workProject, nil
}
func (r *repository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&WorkProject{}, "id = ?", id).Error
}
