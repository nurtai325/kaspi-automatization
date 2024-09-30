package order

import (
	"context"
	"fmt"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	"github.com/nurtai325/kaspi/mailing/internal/models"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

func RefhreshOrders(req kma.GetOrdersRequest, api kma.API) error {
	pages, err := handlePage(req, api)
	if err != nil {
		return err
	}

	errs := make(chan error, pages)

	for i := 1; i < pages; i++ {
		go func() {
			req.PageNumber = i
			_, err = handlePage(req, api)
			errs <- err
		}()
	}

	for i := 1; i < pages; i++ {
		err := <-errs
		if err != nil {
			return err
		}
	}

	return nil
}

func handlePage(req kma.GetOrdersRequest, api kma.API) (int, error) {
	resp, err := api.GetOrders(context.Background(), req)
	if err != nil {
		return 0, err
	}

	err = saveOrders(resp, repositories.Order())
	if err != nil {
		return 0, err
	}

	return resp.Meta.PageCount, nil
}

func saveOrders(resp *kma.OrdersResponse, repo repositories.OrderRepository) error {
	count := len(resp.Data)
	fmt.Print(resp, "\n", count)
	errs := make(chan error)

	for i := 0; i < count; i++ {
		go func(order models.Order) {
			err := repo.Insert(order)
			if err != nil {
				errs <- err
				return
			}
		}(models.Order{
			Id:        resp.Data[i].Attributes.Code,
			Completed: false,
		})
	}

	for i := 0; i < count; i++ {
		err := <-errs
		if err != nil {
			return err
		}
	}

	return nil
}
