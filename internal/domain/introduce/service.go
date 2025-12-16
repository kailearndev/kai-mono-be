package introduce

import (
	"errors"
	"kai-mono-be/pkg/validator"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Service interface {
	ListIntroduces(lang string, limit, offset int) ([]Introduce, int64, error)
	GetIntroduceByID(id uuid.UUID) (Introduce, error)
	CreateIntroduce(p CreateIntroduceDTO) (Introduce, error)
	UpdateIntroduce(id uuid.UUID, dto UpdateIntroduceDTO) (Introduce, error)
	DeleteIntroduce(id uuid.UUID) error
	CountIntroduces() (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListIntroduces(lang string, limit, offset int) ([]Introduce, int64, error) {

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

func (s *service) GetIntroduceByID(id uuid.UUID) (Introduce, error) {

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return Introduce{}, err
	}
	return existing, nil
}

func (s *service) CreateIntroduce(dto CreateIntroduceDTO) (Introduce, error) {

	// validation cơ bản

	// map DTO -> model
	introduce := Introduce{
		Order:     dto.Order,
		IsActive:  dto.IsActive,
		Avatar:    dto.Avatar,
		Facebook:  dto.Facebook,
		Instagram: dto.Instagram,
		Linkedin:  dto.Linkedin,
		Github:    dto.Github,
	}

	if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
		return Introduce{}, err
	}
	// map translations
	for _, t := range dto.Translations {
		if t.Lang == "" || t.Title == "" {
			return Introduce{}, errors.New("ngôn ngữ và tiêu đề bản dịch không được để trống")
		}
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return Introduce{}, err
		}
		introduce.Translations = append(introduce.Translations, IntroduceTranslation{
			Lang:    t.Lang,
			Title:   t.Title,
			Slogan:  t.Slogan,
			Content: pq.StringArray(t.Content),
		})
	}

	if err := s.repo.Create(&introduce); err != nil {
		return Introduce{}, err
	}
	return introduce, nil
}
func (s *service) UpdateIntroduce(id uuid.UUID, dto UpdateIntroduceDTO) (Introduce, error) {
	if len(dto.Translations) > 0 {
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return Introduce{}, err
		}
	}

	var result Introduce
	err := s.repo.WithTx(func(tx *gorm.DB) error {

		// 1️⃣ check tồn tại
		if err := tx.First(&result, "id = ?", id).Error; err != nil {
			return errors.New("introduce not found")
		}

		// 2️⃣ update menu fields (KHÔNG overwrite bừa)
		if dto.Order != nil {
			result.Order = *dto.Order
		}
		if dto.IsActive != nil {
			result.IsActive = *dto.IsActive
		}
		if dto.Avatar != nil {
			result.Avatar = *dto.Avatar
		}
		if dto.Facebook != nil {
			result.Facebook = *dto.Facebook
		}
		if dto.Instagram != nil {
			result.Instagram = *dto.Instagram
		}
		if dto.Linkedin != nil {
			result.Linkedin = *dto.Linkedin
		}
		if dto.Github != nil {
			result.Github = *dto.Github
		}

		if err := tx.Save(&result).Error; err != nil {
			return errors.New("failed to update introduce")
		}

		// 3️⃣ upsert translations
		for _, t := range dto.Translations {

			var mt IntroduceTranslation
			err := tx.
				Where("introduce_id = ? AND lang = ?", id, t.Lang).
				First(&mt).Error

			if errors.Is(err, gorm.ErrRecordNotFound) {
				mt = IntroduceTranslation{
					IntroduceID: id,
					Lang:        t.Lang,
					Title:       t.Title,
					Content:     pq.StringArray(t.Content),
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

func (s *service) DeleteIntroduce(id uuid.UUID) error {
	// optionally: check existence first

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *service) CountIntroduces() (int64, error) {
	return s.repo.Count()
}
