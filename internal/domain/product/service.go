package product

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	ListProducts(limit, offset int) ([]Product, int64, error)
	GetProductByID(id uuid.UUID) (Product, error)
	CreateProduct(p ProductRequest) (Product, error)
	UpdateProduct(id uuid.UUID, input ProductRequest) (Product, error)
	DeleteProduct(id uuid.UUID) error
	CountProducts() (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListProducts(limit, offset int) ([]Product, int64, error) {
	items, err := s.repo.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (s *service) GetProductByID(id uuid.UUID) (Product, error) {

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return Product{}, err
	}
	return existing, nil
}

func (s *service) CreateProduct(req ProductRequest) (Product, error) {
	// validation cơ bản
	if req.Name == "" {
		return Product{}, errors.New("tên sản phẩm không được để trống")
	}

	// map DTO -> model
	product := Product{
		Name:        req.Name,
		SKU:         req.SKU,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURLs:   req.ImageURLs,
		Description: req.Description,
	}

	// gọi repo
	if err := s.repo.Create(&product); err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) UpdateProduct(id uuid.UUID, p ProductRequest) (Product, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return Product{}, err
	}
	existing.Name = p.Name
	existing.Price = p.Price
	existing.Stock = p.Stock
	existing.SKU = p.SKU
	existing.Description = p.Description
	if err := s.repo.Update(&existing); err != nil {
		return Product{}, err
	}
	return existing, nil
}
func (s *service) DeleteProduct(id uuid.UUID) error {
	// optionally: check existence first

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *service) CountProducts() (int64, error) {
	return s.repo.Count()
}
