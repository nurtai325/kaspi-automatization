package repositories

import "github.com/nurtai325/kaspi/mailing/internal/models"

type OrderRepository interface {
	Insert(order models.Order) error
}

type OrderQueueRepository interface {
	Add(id, phone string) error
}
