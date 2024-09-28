package main

import (
	"fmt"
	"sync"

	"github.com/nurtai325/kaspi/mailing/internal/config"
	"github.com/nurtai325/kaspi/mailing/internal/tasks"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	conf, err := config.New()
	fmt.Println(conf)

	var wg sync.WaitGroup
	wg.Add(1)

	stop, err := tasks.Start(conf)
	defer stop()
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
