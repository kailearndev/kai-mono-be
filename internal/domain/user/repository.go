package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(u *User) error
	FindByEmail(email string) (User, error)
	// FindByID(id uuid.UUID) (User, error)
	// UpdateUser(u *User) error
	// DeleteUser(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(u *User) error {
	return r.db.Create(u).Error
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

// func (r *repository) Count() (int64, error) {
// 	var count int64
// 	if err := r.db.Model(&Product{}).Count(&count).Error; err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }
// func (r *repository) FindByID(id uuid.UUID) (Product, error) {
// 	var product Product
// 	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
// 		return Product{}, err
// 	}
// 	return product, nil
// }

// func (r *repository) Create(p *Product) error {
// 	return r.db.Create(p).Error
// }
// func (r *repository) Update(p *Product) error {
// 	return r.db.Save(p).Error
// }
// func (r *repository) Delete(id uuid.UUID) error {
// 	return r.db.Delete(&Product{}, "id = ?", id).Error
// }
