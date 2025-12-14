package home

import (
	"errors"

	"github.com/google/uuid"
)

var ErrDuplicateLocale = errors.New("duplicate locale in translations")

type Service interface {
	ListHome(lang string, limit, offset int) ([]HomeHero, int64, error)
	GetHomeByID(lang string, id uuid.UUID) (HomeHero, error)
	CreateHome(p CreateHomeRequest) (HomeHero, error)
	// UpdateHome(id uuid.UUID, input HeroRequest) (HomeHero, error)
	// DeleteHome(id uuid.UUID) error
	CountHomes() (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListHome(lang string, limit, offset int) ([]HomeHero, int64, error) {
	items, err := s.repo.FindAll(lang, limit, offset)
	if lang == "" {
		lang = "en"
	}
	if err != nil {
		return nil, 0, err
	}
	count, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (s *service) GetHomeByID(lang string, id uuid.UUID) (HomeHero, error) {

	existing, err := s.repo.FindByID(lang, id)
	if err != nil {
		return HomeHero{}, err
	}
	return existing, nil
}

func (s *service) CreateHome(req CreateHomeRequest) (HomeHero, error) {
	// Validate for duplicate locales
	localeMap := make(map[string]bool)
	for _, t := range req.Translations {
		if localeMap[t.Locale] {
			return HomeHero{}, ErrDuplicateLocale
		}
		localeMap[t.Locale] = true
	}

	hero := HomeHero{
		BackgroundImage: req.BackgroundImage,
	}
	for _, t := range req.Translations {
		hero.Translations = append(hero.Translations, HomeHeroTranslation{
			ID:       uuid.New(),
			Locale:   t.Locale,
			Title:    t.Title,
			Subtitle: t.Subtitle,
			CtaText:  t.CtaText,
		})
	}
	// gọi repo
	if err := s.repo.Create(&hero); err != nil {
		return HomeHero{}, err
	}

	return hero, nil
}

// func (s *service) UpdateHome(id uuid.UUID, p HeroRequest) (HomeHero, error) {
// 	existing, err := s.repo.FindByID(id)
// 	if err != nil {
// 		return Product{}, err
// 	}
// 	existing.Name = p.Name
// 	existing.Price = p.Price
// 	existing.Stock = p.Stock
// 	existing.SKU = p.SKU
// 	existing.Description = p.Description
// 	if err := s.repo.Update(&existing); err != nil {
// 		return Product{}, err
// 	}
// 	return existing, nil
// }
// func (s *service) DeleteHome(id uuid.UUID) error {
// 	// optionally: check existence first

// 	_, err := s.repo.FindByID(id)
// 	if err != nil {
// 		return err
// 	}
// 	return s.repo.Delete(id)
// }

func (s *service) CountHomes() (int64, error) {
	return s.repo.Count()
}
