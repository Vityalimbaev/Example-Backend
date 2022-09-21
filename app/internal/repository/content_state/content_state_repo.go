package content_state

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryI interface {
	CreateContentState(*entity.ContentState) (int, error)
	DeleteContentState(*entity.ContentState) error
	GetContentStates() ([]entity.ContentState, error)
	UpdateContentState(*entity.ContentState) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateContentState(input *entity.ContentState) (int, error) {
	stmt, err := r.db.PrepareNamed(`INSERT INTO content_state (title) values (:title) RETURNING id`)
	if err != nil {
		logrus.Error(err)
		return 0, exception.InternalError
	}

	var id int
	if err := stmt.Get(&id, &input); err != nil {
		logrus.Error(err)
		return 0, exception.InternalError
	}

	return id, nil
}

func (r *repository) DeleteContentState(input *entity.ContentState) error {
	query := `DELETE FROM content_state WHERE id=:id`
	if _, err := r.db.NamedExec(query, &input); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}

func (r *repository) GetContentStates() ([]entity.ContentState, error) {
	query := `SELECT id, title FROM content_state`

	var list []entity.ContentState
	if err := r.db.Select(&list, query); err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	return list, nil
}

func (r *repository) UpdateContentState(input *entity.ContentState) error {
	query := `UPDATE content_state SET title=:title WHERE id=:id`
	if _, err := r.db.NamedExec(query, &input); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}
