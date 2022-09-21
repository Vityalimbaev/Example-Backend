package content_action

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryI interface {
	CreateContentAction(*entity.ContentAction) (int, error)
	DeleteContentAction(*entity.ContentAction) error
	GetContentActions() ([]entity.ContentAction, error)
	UpdateContentAction(*entity.ContentAction) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateContentAction(input *entity.ContentAction) (int, error) {
	stmt, err := r.db.PrepareNamed(`INSERT INTO content_action (title) values (:title) RETURNING id`)
	if err != nil {
		logrus.Error(err)
		return 0, exception.InternalError
	}

	var id int
	if err = stmt.Get(&id, &input); err != nil {
		logrus.Error(err)
		return 0, exception.InternalError
	}

	return id, nil
}

func (r *repository) DeleteContentAction(input *entity.ContentAction) error {
	query := `DELETE FROM content_action WHERE id=:id`
	if _, err := r.db.NamedExec(query, &input); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}

func (r *repository) GetContentActions() ([]entity.ContentAction, error) {
	query := `SELECT id, title FROM content_action`

	var list []entity.ContentAction
	if err := r.db.Select(&list, query); err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	return list, nil
}

func (r *repository) UpdateContentAction(input *entity.ContentAction) error {
	query := `UPDATE content_action SET title=:title WHERE id=:id`
	if _, err := r.db.NamedExec(query, &input); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}
