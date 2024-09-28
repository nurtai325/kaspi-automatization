package kaspi

import (
	"context"
	"sync"

	kma "github.com/abdymazhit/kaspi-merchant-api"
)

func RefhreshOrders(api kma.API) error {
	req := getOrderReq()

	resp, err := api.GetOrders(context.Background(), req)
	if err != nil {
		return err
	}

	count := resp.Meta.PageCount

	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		if i == 0 {
			continue
		}
		req.PageNumber = i

		wg.Add(1)
		go func(req kma.GetOrdersRequest) {
			defer wg.Done()
		}(req)
	}

	wg.Wait()

	return nil
}

func getOrderReq() kma.GetOrdersRequest {
	return kma.GetOrdersRequest{}
}

func parseOrder(resp kma.OrdersResponse) {
}
