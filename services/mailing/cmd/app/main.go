package main

import (
	"fmt"
	"sync"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	"github.com/nurtai325/kaspi/mailing/internal/config"
	"github.com/nurtai325/kaspi/mailing/internal/external/kaspi"
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

	api := kma.New(conf.KASPI_TOKEN)
	req := kaspi.GetOrderReq()
	err = kaspi.RefhreshOrders(req, api)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
