package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(limit, offset int) ([]Product, error)
	Count() (int64, error)
	FindByID(id uuid.UUID) (Product, error)
	Create(p *Product) error
	Update(p *Product) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(limit, offset int) ([]Product, error) {
	var products []Product
	if err := r.db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (r *repository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&Product{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (r *repository) FindByID(id uuid.UUID) (Product, error) {
	var product Product
	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) Create(p *Product) error {
	return r.db.Create(p).Error
}
func (r *repository) Update(p *Product) error {
	return r.db.Save(p).Error
}
func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Product{}, "id = ?", id).Error
}
