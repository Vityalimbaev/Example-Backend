package box

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryI interface {
	InsertBox(box *entity.Box) (int, error)
	GetBoxes() ([]entity.Box, error)
	UpdateBox(*entity.Box) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) InsertBox(box *entity.Box) (int, error) {

	stmt, err := r.db.PrepareNamed(`INSERT INTO box (code, content_state_id, unlimited_storage, description) 
													VALUES (:code, :content_state, :unlimited_storage, :description) RETURNING id`)

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	row := stmt.QueryRowx(box)

	var id int

	if err = row.Scan(&id); err != nil {
		logrus.Error(err)
		return 0, err
	}

	_ = stmt.Close()

	return id, nil
}

func (r *repository) GetBoxes() ([]entity.Box, error) {
	query := `SELECT id, code, creation_date, content_state_id, unlimited_storage, description FROM box`

	var boxes []entity.Box

	if err := r.db.Select(&boxes, query); err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	return boxes, nil
}

func (r *repository) UpdateBox(box *entity.Box) error {
	query := `CALL box_update(:code, :content_state_id, :unlimited_storage, :description)`

	if _, err := r.db.Exec(query, box); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}
