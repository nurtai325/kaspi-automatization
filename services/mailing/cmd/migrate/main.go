package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/nurtai325/kaspi/mailing/internal/config"
	"github.com/nurtai325/kaspi/mailing/internal/db"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}
	conf := config.New()

	closeDB, err := db.Connect(conf)
	defer closeDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		fmt.Println("error migrating")
		panic(err)
	}

	fmt.Println("migrated succesfuly")
}
