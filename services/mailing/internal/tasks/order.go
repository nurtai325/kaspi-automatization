package tasks

import (
	"time"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	schedule "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/order"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

func OrdersTask(token string) *schedule.Task {
	return &schedule.Task{
		Interval: time.Duration(order.IntervalMinutes) * time.Minute,
		TaskFunc: func() error {
			api := kma.New(token)
			req := order.GetOrderReq(kma.OrdersStateKaspiDelivery)
			repo := repositories.Order()
			queue := repositories.OrderQueue()
			return order.RefhreshOrders(req, api, repo, queue)
		},
		ErrFunc: errFunc,
	}
}
