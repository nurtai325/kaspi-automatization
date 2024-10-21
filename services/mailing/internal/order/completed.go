package order

import (
	"context"
	"errors"

	kma "github.com/abdymazhit/kaspi-merchant-api"
	"github.com/nurtai325/kaspi/mailing/internal/messaging"
	"github.com/nurtai325/kaspi/mailing/internal/models"
	"github.com/nurtai325/kaspi/mailing/internal/repositories"
)

var (
	ErrOrderNotFound = errors.New("couldn't find the order")
)

func CheckCompleted(
	repo repositories.OrderRepository,
	queue repositories.OrderQueueRepository,
	client models.Client,
) error {
	return queue.Range(func(id string, order models.QueuedOrder) error {
		api := kma.New(order.Token)
		return completeOrder(id, order.ProductCode, order.OrderPhone, repo, queue, api, client)
	})
}

func completeOrder(
	id string,
	productCode string,
	phone string,
	repo repositories.OrderRepository,
	queue repositories.OrderQueueRepository,
	api kma.API,
	client models.Client,
) error {
	orderResp, err := api.GetOrderByCode(context.Background(), id)
	if err != nil {
		return err
	}

	orderData := orderResp.Data
	if len(orderData) < 1 {
		return ErrOrderNotFound
	}

	order := orderData[0]
	status := order.Attributes.Status
	state := order.Attributes.State

	if state != kma.OrdersStateArchive {
		return nil
	} else if status != kma.OrdersStatusCompleted {
		return queue.Remove(id)
	}

	err = repo.Complete(id)
	if err != nil {
		return err
	}

	messenger := messaging.New(client.Id)
	message := messaging.CompletedOrderMessage(
		order.Attributes.Customer.Name,
		order.Attributes.Code,
		productCode,
	)

	err = messenger.Message(models.Message{
		Sender:   client.Phone,
		Receiver: phone,
		Text:     message,
	})
	if err != nil {
		return err
	}

	return queue.Remove(id)
}

func checkStatusCanceled(status string) bool {
	switch status {
	case string(kma.OrdersStatusCancelled):
		return true
	case string(kma.OrdersStatusCancelling):
		return true
	case string(kma.OrdersStatusReturned):
		return true
	case string(kma.OrdersStatusKaspiDeliveryReturnRequested):
		return true
	case string(kma.OrdersStatusReturnAcceptedByMerchant):
		return true
	default:
		return false
	}
}
