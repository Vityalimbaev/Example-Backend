package content_state

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/repository/content_state"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
)

type ServiceI interface {
	CreateContentState(*entity.ContentState) (int, error)
	DeleteContentState(*entity.ContentState) error
	GetContentStates() ([]entity.ContentState, error)
	UpdateContentState(*entity.ContentState) error
}

type service struct {
	contentStateRepo content_state.RepositoryI
}

func NewService(contentStateRepo content_state.RepositoryI) *service {
	return &service{contentStateRepo: contentStateRepo}
}

func (s *service) CreateContentState(input *entity.ContentState) (int, error) {
	if !input.IsValidForSave() {
		return 0, exception.BadRequest
	}
	return s.contentStateRepo.CreateContentState(input)
}

func (s *service) DeleteContentState(input *entity.ContentState) error {
	return s.contentStateRepo.DeleteContentState(input)
}

func (s *service) GetContentStates() ([]entity.ContentState, error) {
	return s.contentStateRepo.GetContentStates()
}

func (s *service) UpdateContentState(input *entity.ContentState) error {
	if !input.IsValidForUpdate() {
		return exception.BadRequest
	}
	return s.contentStateRepo.UpdateContentState(input)
}
