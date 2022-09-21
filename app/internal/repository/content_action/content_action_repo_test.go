package content_action

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Create(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		logrus.Fatal(err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	r := NewRepository(sqlxDB)

	testTable := []struct {
		name      string
		mock      func(input entity.ContentAction)
		input     entity.ContentAction
		expectVal int
		expectErr bool
	}{
		{
			name: "OK",
			input: entity.ContentAction{
				Title: "1234",
			},
			expectVal: 2,
			mock: func(input entity.ContentAction) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(2)
				mock.ExpectQuery("INSERT INTO content_action").
					WithArgs(input.Title).WillReturnRows(rows)
			},
		},
		{
			name:      "empty field",
			input:     entity.ContentAction{},
			expectErr: true,
			mock: func(input entity.ContentAction) {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO content_action").
					WithArgs(input.Title).WillReturnRows(rows)
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input)

			got, err := r.CreateContentAction(&testCase.input)
			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectVal, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		logrus.Fatal(err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	r := NewRepository(sqlxDB)

	testTable := []struct {
		name      string
		mock      func(input entity.ContentAction)
		input     entity.ContentAction
		expectVal int
		expectErr bool
	}{
		{
			name: "OK",
			input: entity.ContentAction{
				Id: 3,
			},
			mock: func(input entity.ContentAction) {
				mock.ExpectExec("DELETE FROM content_action WHERE").WithArgs(input.Id).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name:  "empty field",
			input: entity.ContentAction{},
			mock: func(input entity.ContentAction) {
				mock.ExpectExec("DELETE FROM content_action WHERE").WithArgs(input.Id).WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input)

			err := r.DeleteContentAction(&testCase.input)
			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
