package db

import (
	"database/sql"
	"fmt"
)

type Migrator interface {
	Migrate(*sql.DB) error
	Name() string
}

func Migrate() error {
	migList := []Migrator{
		&Order{},
		&Client{},
		&Customer{},
	}

	db := New()
	for _, migrator := range migList {
		fmt.Printf("running migration: %s\n", migrator.Name())
		err := migrator.Migrate(db)
		if err != nil {
			fmt.Printf("error running migration: %s\n", migrator.Name())
			return err
		}
		fmt.Printf("finished migration: %s\n", migrator.Name())
	}
	return nil
}
