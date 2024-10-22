package tasks

import (
	"time"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	schedule "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/models"
	"github.com/nurtai325/kaspi/mailing/internal/order"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

func NewOrders() *schedule.Task {
	return &schedule.Task{
		Interval: time.Duration(order.IntervalMinutes)*time.Minute + time.Minute,
		TaskFunc: newOrders,
		ErrFunc:  errFunc,
	}
}

func newOrders() error {
	clientRepo := repositories.NewClient()
	clients, err := clientRepo.Get()
	if err != nil {
		return err
	}
	clientsLen := len(clients)
	clients = activeClients(clients, clientsLen)
	clientsLen = len(clients)

	req := order.GetOrderReq(kma.OrdersStateKaspiDelivery)
	repo := repositories.Order()
	queue := repositories.OrderQueue()

	errChan := make(chan error, clientsLen)
	for _, client := range clients {
		go func() {
			api := kma.New(client.Token)
			err := order.RefhreshOrders(req, api, repo, queue, client)
			errChan <- err
		}()
	}

	for i := 0; i < clientsLen; i++ {
		err = <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}

func CompletedOrders() *schedule.Task {
	return &schedule.Task{
		Interval: time.Duration(order.IntervalMinutes)*time.Minute + time.Minute,
		TaskFunc: completedOrders,
		ErrFunc:  errFunc,
	}
}

func completedOrders() error {
	repo := repositories.Order()
	queue := repositories.OrderQueue()
	return order.CheckCompleted(repo, queue, models.Client{})
}

func activeClients(clients []models.Client, length int) []models.Client {
	active := make([]models.Client, 0, length)
	for _, client := range clients {
		now := time.Now().UTC()
		if client.Expires.Valid && now.Before(client.Expires.Time) && client.Connected {
			active = append(active, client)
		} else {
		}
	}
	return active
}
