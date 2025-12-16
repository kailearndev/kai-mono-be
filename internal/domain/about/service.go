package about

import (
	"encoding/json"
	"errors"
	"kai-mono-be/pkg/validator"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	ListAbouts(lang string, limit, offset int) ([]About, int64, error)
	GetAboutByID(id uuid.UUID) (About, error)
	CreateAbout(p CreateAboutDTO) (About, error)
	UpdateAbout(id uuid.UUID, dto UpdateAboutDTO) (About, error)
	DeleteAbout(id uuid.UUID) error
	CountAbouts() (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListAbouts(lang string, limit, offset int) ([]About, int64, error) {

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

func (s *service) GetAboutByID(id uuid.UUID) (About, error) {

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return About{}, err
	}
	return existing, nil
}

func (s *service) CreateAbout(dto CreateAboutDTO) (About, error) {

	// validation cơ bản

	// map DTO -> model
	about := About{
		Order:    dto.Order,
		IsActive: dto.IsActive,
		Avatar:   dto.Avatar,
	}

	// map translations
	for _, t := range dto.Translations {
		if t.Lang == "" {
			return About{}, errors.New("ngôn ngữ và tiêu đề bản dịch không được để trống")
		}
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return About{}, err
		}

		educationsJSON, _ := json.Marshal(t.Educations)
		experiencesJSON, _ := json.Marshal(t.Experiences)
		reviewsJSON, _ := json.Marshal(t.Reviews)

		about.Translations = append(about.Translations, AboutTranslation{
			Lang:        t.Lang,
			Title:       t.Title,
			Content:     t.Content,
			Educations:  educationsJSON,
			Experiences: experiencesJSON,
			Reviews:     reviewsJSON,
		})
	}

	if err := s.repo.Create(&about); err != nil {
		return About{}, err
	}
	return about, nil
}
func (s *service) UpdateAbout(id uuid.UUID, dto UpdateAboutDTO) (About, error) {
	var result About
	if len(dto.Translations) > 0 {
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return About{}, err
		}
	}

	err := s.repo.WithTx(func(tx *gorm.DB) error {

		// 1️⃣ check tồn tại
		if err := tx.First(&result, "id = ?", id).Error; err != nil {
			return errors.New("about not found")
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

		if err := tx.Save(&result).Error; err != nil {
			return errors.New("failed to update about")
		}

		// 3️⃣ upsert translations
		for _, t := range dto.Translations {

			var mt AboutTranslation
			err := tx.
				Where("about_id = ? AND lang = ?", id, t.Lang).
				First(&mt).Error

			experiencesJSON, _ := json.Marshal(t.Experiences)
			educationsJSON, _ := json.Marshal(t.Educations)

			reviewsJSON, _ := json.Marshal(t.Reviews)

			if errors.Is(err, gorm.ErrRecordNotFound) {
				mt = AboutTranslation{
					AboutID:     id,
					Lang:        t.Lang,
					Title:       t.Title,
					Content:     t.Content,
					Educations:  educationsJSON,
					Experiences: experiencesJSON,
					Reviews:     reviewsJSON,
				}
				if err := tx.Create(&mt).Error; err != nil {
					return err
				}
			} else if err == nil {
				// ✏️ update
				mt.Title = t.Title
				mt.Content = t.Content
				mt.Educations = educationsJSON
				mt.Experiences = experiencesJSON
				mt.Reviews = reviewsJSON

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

func (s *service) DeleteAbout(id uuid.UUID) error {
	// optionally: check existence first

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *service) CountAbouts() (int64, error) {
	return s.repo.Count()
}
