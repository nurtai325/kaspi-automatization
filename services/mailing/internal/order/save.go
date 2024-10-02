package order

import (
	"errors"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	"github.com/nurtai325/kaspi/mailing/internal/models"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

func saveOrders(
	resp *kma.OrdersResponse,
	repo repositories.OrderRepository,
	queue repositories.OrderQueueRepository,
) error {
	count := len(resp.Data)
	errs := make(chan error)

	for i := 0; i < count; i++ {
		go save(models.Order{
			Id:        resp.Data[i].Attributes.Code,
			Completed: false,
		}, errs, repo, queue)
	}

	for i := 0; i < count; i++ {
		err := <-errs
		if err != nil {
			return err
		}
	}

	return nil
}

func save(
	order models.Order,
	errs chan error,
	repo repositories.OrderRepository,
	queue repositories.OrderQueueRepository,
) {
	err := repo.Insert(order)
	if err == nil {
		err = queue.Add(order.Id, order.Phone)
		if err != nil {
			errs <- err
			return
		}
		return
	}

	if errors.Is(err, repositories.ErrRecordIsPresent) {
		err = queue.Add(order.Id, order.Phone)
		if err != nil {
			errs <- err
			return
		}
		return
	}
	errs <- err
	return
}
