package content_action

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/repository/content_action"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
)

type ServiceI interface {
	CreateContentAction(*entity.ContentAction) (int, error)
	DeleteContentAction(*entity.ContentAction) error
	GetContentActions() ([]entity.ContentAction, error)
	UpdateContentAction(*entity.ContentAction) error
}

type service struct {
	contentActionRepo content_action.RepositoryI
}

func NewService(contentActionRepo content_action.RepositoryI) *service {
	return &service{contentActionRepo: contentActionRepo}
}

func (s *service) CreateContentAction(input *entity.ContentAction) (int, error) {
	if !input.IsValidForSave() {
		return 0, exception.BadRequest
	}
	return s.contentActionRepo.CreateContentAction(input)
}

func (s *service) DeleteContentAction(input *entity.ContentAction) error {
	return s.contentActionRepo.DeleteContentAction(input)
}

func (s *service) GetContentActions() ([]entity.ContentAction, error) {
	return s.contentActionRepo.GetContentActions()
}

func (s *service) UpdateContentAction(input *entity.ContentAction) error {
	if !input.IsValidForUpdate() {
		return exception.BadRequest
	}
	return s.contentActionRepo.UpdateContentAction(input)
}
