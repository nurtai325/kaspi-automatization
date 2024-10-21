package order

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	"github.com/nurtai325/kaspi/mailing/internal/messaging"
	"github.com/nurtai325/kaspi/mailing/internal/models"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

var (
	ErrIncorrectData = errors.New("request doesn't contain valid data")
)

func saveOrders(
	resp *kma.OrdersResponse,
	repo repositories.OrderRepository,
	queue repositories.OrderQueueRepository,
	api kma.API,
	client models.Client,
) error {
	count := len(resp.Data)
	errs := make(chan error)

	for i := 0; i < count; i++ {
		order := resp.Data[i]
		attributes := order.Attributes

		go save(models.Order{
			Id:        attributes.Code,
			Completed: false,
			Phone:     attributes.Customer.CellPhone,
			Sum:       int64(attributes.TotalPrice),
			Customer:  attributes.Customer.Name,
			KaspiId:   order.Id,
		}, errs, repo, queue, api, client)
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
	client models.Client,
) {
	entryResp, err := api.GetOrderEntries(context.Background(), order.KaspiId)
	if err != nil {
		errs <- err
		return
	}
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

	phone := "7" + order.Phone
	err = queue.Add(order.Id, models.QueuedOrder{
		ClientId:    client.Id,
		Token:       client.Token,
		ProductCode: order.ProductCode,
		ClientPhone: client.Phone,
		OrderPhone:  phone,
	})
	if err != nil {
		errs <- err
		return
	}

	exists, err := repo.Exists(order.Id)
	if err != nil {
		errs <- err
		return
	}
	if exists {
		errs <- nil
		return
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

	messenger := messaging.New(client.Id)
	message := messaging.NewOrderMessage(order.Customer, order.Id, entries)
	err = messenger.Message(models.Message{
		Sender:   client.Phone,
		Receiver: phone,
		Text:     message,
	})
	if err != nil {
		errs <- err
		return
	}

	errs <- nil
	return
}
