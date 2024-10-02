package main

import (
	"sync"

	_ "github.com/lib/pq"
	scheduling "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/config"
	"github.com/nurtai325/kaspi/mailing/internal/db"
	"github.com/nurtai325/kaspi/mailing/internal/models"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
	"github.com/nurtai325/kaspi/mailing/internal/tasks"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}
	conf, err := config.New()

	var wg sync.WaitGroup
	wg.Add(1)

	jobs := []*scheduling.Task{tasks.OrdersTask(conf.KASPI_TOKEN)}
	stop, err := tasks.Start(conf, jobs)
	defer stop()
	if err != nil {
		panic(err)
	}

	closeDB, err := db.Connect(conf)
	defer closeDB()
	if err != nil {
		panic(err)
	}

	repo := repositories.Order()
	err = repo.Insert(models.Order{
		Id:        "yessss",
		Completed: true,
	})
	panic(err)
	wg.Done()

	wg.Wait()
}
