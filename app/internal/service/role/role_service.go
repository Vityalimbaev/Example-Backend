package role

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/repository/role"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
)

type ServiceI interface {
	CreateRole(*entity.Role) (int, error)
	DeleteRole(*entity.Role) error
	GetRoles() ([]entity.Role, error)
	UpdateRole(*entity.Role) error
}

type service struct {
	roleRepo role.RepositoryI
}

func NewService(roleRepo role.RepositoryI) *service {
	return &service{roleRepo: roleRepo}
}

func (s *service) CreateRole(input *entity.Role) (int, error) {
	if !input.IsValidForSave() {
		return 0, exception.BadRequest
	}
	return s.roleRepo.CreateRole(input)
}

func (s *service) DeleteRole(input *entity.Role) error {
	return s.roleRepo.DeleteRole(input)
}

func (s *service) GetRoles() ([]entity.Role, error) {
	return s.roleRepo.GetRoles()
}

func (s *service) UpdateRole(input *entity.Role) error {
	if !input.IsValidForUpdate() {
		return exception.BadRequest
	}
	return s.roleRepo.UpdateRole(input)
}
