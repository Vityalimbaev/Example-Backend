package history

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/repository/history"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
)

type ServiceI interface {
	CreateActionHistory(*entity.ActionHistory) (int, error)
	GetActionHistory() ([]entity.ActionHistory, error)
}

type service struct {
	historyRepo history.RepositoryI
}

func NewService(historyRepo history.RepositoryI) *service {
	return &service{historyRepo: historyRepo}
}

func (s *service) CreateActionHistory(ActionHistoryArg *entity.ActionHistory) (int, error) {
	if !ActionHistoryArg.IsValidForSave() {
		return 0, exception.BadRequest
	}
	return s.historyRepo.InsertActionHistory(ActionHistoryArg)
}

func (s *service) GetActionHistory() ([]entity.ActionHistory, error) {
	return s.historyRepo.GetActionHistory()
}
