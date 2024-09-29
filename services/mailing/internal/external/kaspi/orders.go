package kaspi

import (
	"sync"

	kma "github.com/abdymazhit/kaspi-merchant-api"
)

func RefhreshOrders(req kma.GetOrdersRequest, api kma.API) error {
	pages, err := handlePage(req, api)

	errs := []error{}
	var wg sync.WaitGroup

	for i := 1; i < pages; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req.PageNumber = i
			_, err = handlePage(req, api)
			errs = append(errs, err)
		}()
	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func handlePage(req kma.GetOrdersRequest, api kma.API) (int, error) {
	return 1, nil
}

func GetOrderReq() kma.GetOrdersRequest {
	return kma.GetOrdersRequest{}
}

func saveOrders(resp *kma.OrdersResponse) error {
	return nil
}
