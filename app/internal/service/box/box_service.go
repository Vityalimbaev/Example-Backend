package box

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/repository/box"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
)

type ServiceI interface {
	CreateBox(*entity.Box) (int, error)
	GetBoxes() ([]entity.Box, error)
	UpdateBox(*entity.Box) error
}

type service struct {
	BoxRepo box.RepositoryI
}

func NewService(BoxRepo box.RepositoryI) *service {
	return &service{BoxRepo: BoxRepo}
}

func (s *service) CreateBox(boxArg *entity.Box) (int, error) {
	if !boxArg.IsValidForSave() {
		return 0, exception.BadRequest
	}
	return s.BoxRepo.InsertBox(boxArg)
}

func (s *service) GetBoxes() ([]entity.Box, error) {
	return s.BoxRepo.GetBoxes()
}

func (s *service) UpdateBox(input *entity.Box) error {
	if !input.IsValidForUpdate() {
		return exception.BadRequest
	}
	return s.BoxRepo.UpdateBox(input)
}
