package db

import (
	"fmt"
)

type Migrator interface {
	Migrate() error
	Name() string
}

func Migrate() error {
	migList := []Migrator{Order{}}

	for _, migrator := range migList {
		fmt.Printf("running migration: %s\n", migrator.Name())
		err := migrator.Migrate()
		if err != nil {
			fmt.Printf("error running migration: %s\n", migrator.Name())
			return err
		}
		fmt.Printf("finished migration: %s\n", migrator.Name())
	}
	return nil
}
