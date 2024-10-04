package order

import (
	"context"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

func RefhreshOrders(
	req kma.GetOrdersRequest,
	api kma.API,
	repo repositories.OrderRepository,
	queue repositories.OrderQueueRepository,
) error {
	pages, err := handleOrderPage(req, api, repo, queue)
	if err != nil {
		return err
	}

	errs := make(chan error, pages)

	for i := 1; i < pages; i++ {
		go func(errs chan error) {
			req.PageNumber = i
			_, err = handleOrderPage(req, api, repo, queue)
			errs <- err
		}(errs)
	}

	for i := 1; i < pages; i++ {
		err := <-errs
		if err != nil {
			return err
		}
	}

	return nil
}

func handleOrderPage(
	req kma.GetOrdersRequest,
	api kma.API,
	repo repositories.OrderRepository,
	queue repositories.OrderQueueRepository,
) (int, error) {
	resp, err := api.GetOrders(context.Background(), req)
	if err != nil {
		return 0, err
	}

	err = saveOrders(resp, repo, queue, api)
	if err != nil {
		return 0, err
	}

	return resp.Meta.PageCount, nil
}
