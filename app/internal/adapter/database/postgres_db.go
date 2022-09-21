package database

import (
	"fmt"
	"github.com/Vityalimbaev/Example-Backend/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func GetDbConnection(dbConfig *config.DbConfig) *sqlx.DB {

	cmd := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.SSLMode)

	logrus.Info(dbConfig)

	db, err := sqlx.Open("pgx", cmd)

	if err != nil || db == nil {
		logrus.Panic("Failed init  DB ", err)
		return nil
	}

	start := time.Now()
	for err = db.Ping(); err != nil; {
		logrus.Error("Failed init database ", err)

		if time.Now().After(start.Add(30 * time.Second)) {
			logrus.Panic("Failed init database")
		}

		time.Sleep(5 * time.Second)
	}

	return db
}

func UpDBMigrations(db *sqlx.DB) {
	if err := goose.SetDialect("pgx"); err != nil {
		panic(err)
	}

	f, _ := os.Getwd()
	migrationsDir := strings.Split(f, "/app")[0] + "/db/migrations"

	if err := goose.Up(db.DB, migrationsDir); err != nil {
		panic(err)
	}
}
