package history

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryI interface {
	InsertActionHistory(ActionHistory *entity.ActionHistory) (int, error)
	GetActionHistory() ([]entity.ActionHistory, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) InsertActionHistory(ActionHistory *entity.ActionHistory) (int, error) {
	var id int

	stmt, err := r.db.PrepareNamed(`INSERT INTO history (content_action_id, box_id, record_id, description, account_id) 
													VALUES (:content_action_id, :box_id, :record_id, :description, :account_id) RETURNING id`)

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	row := stmt.QueryRowx(ActionHistory)

	if err = row.Scan(&id); err != nil {
		logrus.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *repository) GetActionHistory() ([]entity.ActionHistory, error) {
	query := `SELECT id, content_action_id, box_id, record_id, description, account_id, datetime FROM history`

	var ActionHistoryList []entity.ActionHistory

	if err := r.db.Select(&ActionHistoryList, query); err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	return ActionHistoryList, nil
}
