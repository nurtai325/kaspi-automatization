package order

import (
	"context"
	"encoding/json"
	"strings"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	"github.com/nurtai325/kaspi/mailing/internal/messaging"
	"github.com/nurtai325/kaspi/mailing/internal/models"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

var (
	ErrIncorrectData = "request doesn't contain valid data"
)

func saveOrders(
	resp *kma.OrdersResponse,
	repo repositories.OrderRepository,
	queue repositories.OrderQueueRepository,
	api kma.API,
) error {
	count := len(resp.Data)
	errs := make(chan error)

	for i := 0; i < count; i++ {
		attributes := resp.Data[i].Attributes

		go save(models.Order{
			Id:        attributes.Code,
			Completed: false,
			Phone:     attributes.Customer.CellPhone,
			Sum:       int64(attributes.TotalPrice),
			Customer:  attributes.Customer.Name,
		}, errs, repo, queue, api)
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
	api kma.API,
) {
	entryResp, err := api.GetOrderEntries(context.Background(), order.Id)
	entries := make([]models.Entry, len(entryResp.Data))

	for i, entry := range entryResp.Data {
		if i < 1 {
			substrings := strings.Split(entry.Attributes.Offer.Code, "_")
			if len(substrings) != 2 {
				errs <- err
				return
			}
			order.ProductCode = substrings[0]
		}

		entries[i] = models.Entry{
			Id:           entry.Id,
			Price:        int64(entry.Attributes.BasePrice),
			DeliveryCost: int64(entry.Attributes.DeliveryCost),
			Quantity:     entry.Attributes.Quantity,
			ProductName:  entry.Attributes.Offer.Name,
			ProductCode:  order.ProductCode,
		}
	}

	entriesJson, err := json.Marshal(entries)
	if err != nil {
		errs <- err
		return
	}
	order.Entries = string(entriesJson)

	err = repo.Insert(order)
	if err != nil {
		errs <- err
		return
	}

	messenger := messaging.New()
	err = messenger.Message(order.Phone, "")
	if err != nil {
		errs <- err
		return
	}

	err = queue.Add(order.Id, order.ProductCode)
	if err != nil {
		errs <- err
		return
	}
	return
}
