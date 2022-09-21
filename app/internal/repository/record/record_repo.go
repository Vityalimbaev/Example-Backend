package record

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryI interface {
	InsertRecord(*entity.Record) (int, error)
	GetRecords(params *entity.RecordSearchParams) ([]entity.Record, error)
	UpdateRecord(*entity.Record) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) InsertRecord(input *entity.Record) (int, error) {

	stmt, err := r.db.PrepareNamed(`INSERT INTO record (archived_date, branch, pcode, last_treat, content_state_id, box_id)
			values (to_timestamp(:archived_date), :branch, :pcode, to_timestamp(:last_treat), :content_state_id, :box_id) RETURNING id`)
	if err != nil {
		logrus.Error(err)
		return 0, exception.InternalError
	}

	var id int
	err = stmt.Get(&id, input)
	if err != nil {
		logrus.Error(err)
		return 0, exception.InternalError
	}

	return id, nil
}

func (r *repository) GetRecords(params *entity.RecordSearchParams) ([]entity.Record, error) {
	logrus.Debug(params)

	stmt, err := r.db.PrepareNamed(`SELECT * FROM record_select(:id, :pcode, :branch, :box_id, :content_state_id, 
   				 :start_archived_date, :end_archived_date,:start_creation_date,
    				:end_creation_date, :start_last_treat, :end_last_treat)`)

	if err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	var list []entity.Record

	if err = stmt.Select(&list, params); err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	return list, nil
}

func (r *repository) UpdateRecord(input *entity.Record) error {
	query := `CALL record_update(:id, to_timestamp(:archived_date), :branch, :pcode, to_timestamp(:last_treat))`
	if _, err := r.db.NamedExec(query, &input); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}
