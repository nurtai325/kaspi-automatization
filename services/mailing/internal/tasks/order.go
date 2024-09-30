package tasks

import (
	"time"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	schedule "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/external/kaspi/order"
)

func orderTask(token string) *schedule.Task {
	return &schedule.Task{
		Interval: time.Duration(order.IntervalMinutes) * time.Minute,
		TaskFunc: func() error {
			api := kma.New(token)
			req := order.GetOrderReq(kma.OrdersStateKaspiDelivery)
			return order.RefhreshOrders(req, api)
		},
		ErrFunc: errFunc,
	}
}
