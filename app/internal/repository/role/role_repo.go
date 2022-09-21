package role

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryI interface {
	CreateRole(*entity.Role) (int, error)
	DeleteRole(*entity.Role) error
	GetRoles() ([]entity.Role, error)
	UpdateRole(*entity.Role) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateRole(input *entity.Role) (int, error) {
	stmt, err := r.db.PrepareNamed(`INSERT INTO role (title) values (:title) RETURNING id`)
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

func (r *repository) DeleteRole(input *entity.Role) error {
	query := `DELETE FROM role WHERE id=:id`
	if _, err := r.db.NamedExec(query, &input); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}

func (r *repository) GetRoles() ([]entity.Role, error) {
	query := `SELECT id, title FROM role`

	var list []entity.Role
	if err := r.db.Select(&list, query); err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	return list, nil
}

func (r *repository) UpdateRole(input *entity.Role) error {
	query := `UPDATE role SET title=:title WHERE id=:id`
	if _, err := r.db.NamedExec(query, &input); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}
