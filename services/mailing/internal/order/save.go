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
) error {
	count := len(resp.Data)
	errs := make(chan error)

	save := func(order models.Order, errs chan error) {
		err := repo.Insert(order)
		if err != nil {
			if !errors.Is(err, repositories.ErrRecordIsPresent) {
				errs <- err
			}
			return
		}
	}

	for i := 0; i < count; i++ {
		go save(models.Order{
			Id:        resp.Data[i].Attributes.Code,
			Completed: false,
		}, errs)
	}

	for i := 0; i < count; i++ {
		err := <-errs
		if err != nil {
			return err
		}
	}

	return nil
}
