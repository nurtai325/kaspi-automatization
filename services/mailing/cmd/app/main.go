package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	kaspi_merchant "github.com/abdymazhit/kaspi-merchant-api"
	_ "github.com/lib/pq"
	scheduling "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/config"
	"github.com/nurtai325/kaspi/mailing/internal/db"
	"github.com/nurtai325/kaspi/mailing/internal/order"
	"github.com/nurtai325/kaspi/mailing/internal/tasks"
)

// TODO: use ctx.Context instead of buffered channels
// TODO: adding logging
// TODO: finish clients api
func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}
	conf, err := config.New()

	var wg sync.WaitGroup
	wg.Add(1)

	jobs := []*scheduling.Task{tasks.NewOrders(conf.KASPI_TOKEN)}
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

	api := kaspi_merchant.New(conf.KASPI_TOKEN)
	resp, err := api.GetOrders(context.Background(), order.GetOrderReq(kaspi_merchant.OrdersStateArchive))
	if err != nil {
		panic(err)
	}

	js, _ := json.Marshal(resp)
	fmt.Println(string(js))

	wg.Done()

	wg.Wait()
}
