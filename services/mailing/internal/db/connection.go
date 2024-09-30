package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nurtai325/kaspi/mailing/internal/models"
)

var db *sql.DB

const (
	driver = "postgres"
)

type close func() error

func Connect(conf models.Config) (close, error) {
	port, err := strconv.Atoi(conf.DB_PORT)
	if err != nil {
		return func() error { return nil }, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, port, conf.DB_USER, conf.DB_PASSWORD, conf.DB_NAME)

	pgDB, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return func() error { return nil }, err
	}

	err = pgDB.Ping()
	if err != nil {
		return func() error { return nil }, err
	}

	db = pgDB
	return pgDB.Close, nil
}

func GetDBConnection() *sql.DB {
	return db
}
