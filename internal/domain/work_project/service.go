package work_project

import (
	"errors"
	"kai-mono-be/pkg/validator"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	ListWorkProjects(lang string, limit, offset int) ([]WorkProject, int64, error)
	GetWorkProjectByID(id uuid.UUID) (WorkProject, error)
	GetWorkProjectBySlug(slug string) (WorkProject, error)
	CreateWorkProject(p CreateWorkProjectDTO) (WorkProject, error)
	UpdateWorkProject(id uuid.UUID, dto UpdateWorkProjectDTO) (WorkProject, error)
	DeleteWorkProject(id uuid.UUID) error
	CountWorkProjects() (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListWorkProjects(lang string, limit, offset int) ([]WorkProject, int64, error) {

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

func (s *service) GetWorkProjectBySlug(slug string) (WorkProject, error) {

	existing, err := s.repo.FindBySlug(slug)
	if err != nil {
		return WorkProject{}, err
	}
	return existing, nil
}

func (s *service) GetWorkProjectByID(id uuid.UUID) (WorkProject, error) {

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return WorkProject{}, err
	}
	return existing, nil
}

func (s *service) CreateWorkProject(dto CreateWorkProjectDTO) (WorkProject, error) {

	// validation cơ bản

	// map DTO -> model
	workProject := WorkProject{
		Slug:      dto.Slug,
		Order:     dto.Order,
		IsActive:  dto.IsActive,
		Thumbnail: dto.Thumbnail,
		Link:      dto.Link,
	}

	// map translations
	for _, t := range dto.Translations {
		if t.Lang == "" {
			return WorkProject{}, errors.New("ngôn ngữ và tiêu đề bản dịch không được để trống")
		}
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return WorkProject{}, err
		}
		workProject.Translations = append(workProject.Translations, WorkProjectTranslation{
			Lang:         t.Lang,
			Title:        t.Title,
			Slogan:       t.Slogan,
			Content:      t.Content,
			Subtitle:     t.Subtitle,
			Duration:     t.Duration,
			ClientName:   t.ClientName,
			Services:     t.Services,
			Technologies: t.Technologies,
			ProjectUrl:   t.ProjectUrl,
			ProjectName:  t.ProjectName,
		})
	}

	if err := s.repo.Create(&workProject); err != nil {
		return WorkProject{}, err
	}
	return workProject, nil
}
func (s *service) UpdateWorkProject(id uuid.UUID, dto UpdateWorkProjectDTO) (WorkProject, error) {
	if len(dto.Translations) > 0 {
		if err := validator.ValidateUniqueLang(dto.Translations); err != nil {
			return WorkProject{}, err
		}
	}

	var result WorkProject
	err := s.repo.WithTx(func(tx *gorm.DB) error {

		// 1️⃣ check tồn tại
		if err := tx.First(&result, "id = ?", id).Error; err != nil {
			return errors.New("work project not found")
		}

		// 2️⃣ update menu fields (KHÔNG overwrite bừa)
		if dto.Order != nil {
			result.Order = *dto.Order
		}
		if dto.IsActive != nil {
			result.IsActive = *dto.IsActive
		}
		if dto.Thumbnail != nil {
			result.Thumbnail = *dto.Thumbnail
		}
		if dto.Link != nil {
			result.Link = *dto.Link
		}

		if err := tx.Save(&result).Error; err != nil {
			return errors.New("failed to update introduce")
		}

		// 3️⃣ upsert translations
		for _, t := range dto.Translations {

			var mt WorkProjectTranslation
			err := tx.
				Where("work_project_id = ? AND lang = ?", id, t.Lang).
				First(&mt).Error

			if errors.Is(err, gorm.ErrRecordNotFound) {
				mt = WorkProjectTranslation{
					WorkProjectID: id,
					Lang:          t.Lang,
					Title:         t.Title,
					Content:       t.Content,

					Slogan: t.Slogan,

					Subtitle:     t.Subtitle,
					Duration:     t.Duration,
					ClientName:   t.ClientName,
					Services:     t.Services,
					Technologies: t.Technologies,
					ProjectUrl:   t.ProjectUrl,
					ProjectName:  t.ProjectName,
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

func (s *service) DeleteWorkProject(id uuid.UUID) error {
	// optionally: check existence first

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *service) CountWorkProjects() (int64, error) {
	return s.repo.Count()
}
