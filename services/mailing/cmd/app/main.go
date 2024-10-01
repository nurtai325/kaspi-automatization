package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	_ "github.com/lib/pq"
	scheduling "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/config"
	"github.com/nurtai325/kaspi/mailing/internal/db"
	"github.com/nurtai325/kaspi/mailing/internal/order"
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

	api := kma.New(conf.KASPI_TOKEN)
	req := order.GetOrderReq(kma.OrdersStateKaspiDelivery)
	resp, err := api.GetOrders(context.Background(), req)
	if err != nil {
		panic(err)
	}

	orders, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(orders))
	wg.Done()

	wg.Wait()
}
