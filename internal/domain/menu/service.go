package menu

import (
	"errors"
	"kai-mono-be/pkg/validator"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	ListMenus(lang string, limit, offset int) ([]Menu, int64, error)
	GetMenuByID(id uuid.UUID) (Menu, error)
	CreateMenu(p CreateMenuDTO) (Menu, error)
	UpdateMenu(id uuid.UUID, dto UpdateMenuDTO) (Menu, error)
	DeleteMenu(id uuid.UUID) error
	CountMenus() (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListMenus(lang string, limit, offset int) ([]Menu, int64, error) {

	items, err := s.repo.FindAll(lang, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (s *service) GetMenuByID(id uuid.UUID) (Menu, error) {

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return Menu{}, err
	}
	return existing, nil
}

func (s *service) CreateMenu(dto CreateMenuDTO) (Menu, error) {
	if len(dto.Translations) > 0 {
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return Menu{}, err
		}
	}
	// validation cơ bản
	if dto.Slug == "" {
		return Menu{}, errors.New("tên slug không được để trống")
	}
	// map DTO -> model
	menu := Menu{
		Slug:     dto.Slug,
		Order:    dto.Order,
		IsActive: dto.IsActive,
	}
	isSlugExist, _ := s.repo.FindBySlug(dto.Slug)
	if isSlugExist.ID != uuid.Nil {
		return Menu{}, errors.New("slug đã tồn tại")
	}
	if len(dto.Translations) == 0 {
		return Menu{}, errors.New("cần ít nhất một bản dịch cho menu")
	}
	// map translations
	for _, t := range dto.Translations {
		if t.Lang == "" || t.Title == "" {
			return Menu{}, errors.New("ngôn ngữ và tiêu đề bản dịch không được để trống")
		}
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return Menu{}, err
		}
		menu.Translations = append(menu.Translations, MenuTranslation{
			Lang:  t.Lang,
			Title: t.Title,
		})
	}

	if err := s.repo.Create(&menu); err != nil {
		return Menu{}, err
	}
	return menu, nil
}
func (s *service) UpdateMenu(id uuid.UUID, dto UpdateMenuDTO) (Menu, error) {
	if len(dto.Translations) > 0 {
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return Menu{}, err
		}
	}

	var result Menu

	err := s.repo.WithTx(func(tx *gorm.DB) error {

		// 1️⃣ check tồn tại
		if err := tx.First(&result, "id = ?", id).Error; err != nil {
			return errors.New("menu not found")
		}

		// 2️⃣ update menu fields (KHÔNG overwrite bừa)
		if dto.Order != nil {
			result.Order = *dto.Order
		}
		if dto.IsActive != nil {
			result.IsActive = *dto.IsActive
		}
		if dto.Slug != nil {
			result.Slug = *dto.Slug
		}

		if err := tx.Save(&result).Error; err != nil {
			return errors.New("failed to update menu")
		}

		// 3️⃣ upsert translations
		for _, t := range dto.Translations {

			var mt MenuTranslation
			err := tx.
				Where("menu_id = ? AND lang = ?", id, t.Lang).
				First(&mt).Error

			if errors.Is(err, gorm.ErrRecordNotFound) {
				mt = MenuTranslation{
					MenuID: id,
					Lang:   t.Lang,
					Title:  t.Title,
				}
				if err := tx.Create(&mt).Error; err != nil {
					return err
				}
			} else if err == nil {
				mt.Title = t.Title
				if err := tx.Save(&mt).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}

		return nil
	})

	return result, err
}

func (s *service) DeleteMenu(id uuid.UUID) error {
	// optionally: check existence first

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *service) CountMenus() (int64, error) {
	return s.repo.Count()
}
