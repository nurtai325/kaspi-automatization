package order

import (
	"time"

	kma "github.com/abdymazhit/kaspi-merchant-api"
)

const (
	pageSize        int = 15
	IntervalMinutes int = 3
)

func GetOrderReq(state kma.OrdersState) kma.GetOrdersRequest {
	return kma.GetOrdersRequest{
		PageSize:   pageSize,
		PageNumber: 0,
		Filter: struct {
			Orders struct {
				CreationDateGe    time.Time
				CreationDateLe    time.Time
				State             kma.OrdersState
				Status            kma.OrdersStatus
				DeliveryType      kma.OrdersDeliveryType
				SignatureRequired bool
			}
		}{
			Orders: struct {
				CreationDateGe    time.Time
				CreationDateLe    time.Time
				State             kma.OrdersState
				Status            kma.OrdersStatus
				DeliveryType      kma.OrdersDeliveryType
				SignatureRequired bool
			}{
				CreationDateGe: time.Now().AddDate(0, 0, -IntervalMinutes),
				CreationDateLe: time.Now(),
				State:          state,
			},
		},
	}
}
