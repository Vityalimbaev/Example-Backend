package user

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryI interface {
	InsertUser(user *entity.User) (int, error)
	UpdateUser(user *entity.User) error
	GetUsers(userSearchParams *entity.UserSearchParams) ([]entity.User, error)

	GetUserSession(userId int) (*entity.UserSession, error)
	InsertUserSession(userSession *entity.UserSession) error
	UpdateUserSession(userSession *entity.UserSession) error
	DeleteUserSession(userId int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) InsertUser(user *entity.User) (int, error) {
	query := `INSERT INTO account (name, email, username, password, branch, role_id, active_status) VALUES (:name, :email, :username, :password, :branch, :role_id, true) RETURNING id`

	row, err := r.db.NamedQuery(query, &user)

	if err != nil {
		logrus.Error(err)
		return 0, exception.InternalError
	}

	defer func(row *sqlx.Rows) {
		_ = row.Close()
	}(row)

	var id int

	row.Next()

	if err = row.Scan(&id); err != nil {
		logrus.Error(err)
		return 0, exception.InternalError
	}

	return id, err
}

func (r *repository) UpdateUser(user *entity.User) error {
	query := `CALL user_update (:id, :name, :username, :email, :password, :branch, :role_id)`

	if _, err := r.db.NamedExec(query, &user); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}

func (r *repository) GetUsers(userSearchParams *entity.UserSearchParams) ([]entity.User, error) {
	query, err := r.db.PrepareNamed(`SELECT * FROM user_select (:id, :name, :username,:email, :role_id, :branch)`)

	if err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	rows, err := query.Queryx(userSearchParams)

	defer func(rows *sqlx.Rows) {
		if rows != nil {
			_ = rows.Close()
		}
	}(rows)

	if err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	var result []entity.User

	for rows.Next() {
		user := entity.User{}
		if err = rows.StructScan(&user); err != nil {
			logrus.Error(err)
			return nil, exception.InternalError
		}

		result = append(result, user)
	}

	return result, nil
}

func (r *repository) GetUserSession(userId int) (*entity.UserSession, error) {
	query := `SELECT id, refresh_token, date_expire, account_id FROM user_session WHERE account_id = $1`

	var result []entity.UserSession
	err := r.db.Select(&result, query, userId)

	if err != nil {
		logrus.Error(err)
		return nil, exception.InternalError
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func (r *repository) InsertUserSession(userSession *entity.UserSession) error {
	query := `INSERT INTO user_session (refresh_token, date_expire, account_id) VALUES (:refresh_token, :date_expire, :account_id)
						ON CONFLICT (account_id) DO UPDATE SET refresh_token = excluded.refresh_token , date_expire = excluded.date_expire`

	row, err := r.db.NamedQuery(query, &userSession)

	if err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	defer func(row *sqlx.Rows) {
		_ = row.Close()
	}(row)

	return err
}

func (r *repository) UpdateUserSession(userSession *entity.UserSession) error {
	query := `UPDATE user_session SET refresh_token = :refresh_token, date_expire = :date_expire WHERE account_id = $1`

	if _, err := r.db.NamedExec(query, userSession.UserId); err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return nil
}

func (r *repository) DeleteUserSession(userId int) error {
	query := "DELETE FROM user_session WHERE account_id = $1 "
	_, err := r.db.Exec(query, userId)

	if err != nil {
		logrus.Error(err)
		return exception.InternalError
	}

	return err
}
