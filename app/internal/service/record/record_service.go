package record

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/repository/record"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"time"
)

type ServiceI interface {
	CreateRecord(*entity.Record) (int, error)
	GetRecords(params *entity.RecordSearchParams) ([]entity.Record, error)
	UpdateRecord(*entity.Record) error
}

type service struct {
	recordRepo record.RepositoryI
}

func NewService(recordRepo record.RepositoryI) *service {
	return &service{recordRepo: recordRepo}
}

func (s *service) CreateRecord(input *entity.Record) (int, error) {
	if !input.IsValidForSave() {
		return 0, exception.BadRequest
	}
	return s.recordRepo.InsertRecord(input)
}

func (s *service) GetRecords(params *entity.RecordSearchParams) ([]entity.Record, error) {

	if params.EndCreationDate == 0 {
		params.EndCreationDate = time.Now().Unix()
	}

	if params.EndArchivedDate == 0 {
		params.EndArchivedDate = time.Now().Unix()
	}

	if params.EndLastTreat == 0 {
		params.EndLastTreat = time.Now().Unix()
	}

	return s.recordRepo.GetRecords(params)
}

func (s *service) UpdateRecord(input *entity.Record) error {
	if !input.IsValidForUpdate() {
		return exception.BadRequest
	}
	return s.recordRepo.UpdateRecord(input)
}
