package repositories

import "github.com/nurtai325/kaspi/mailing/internal/models"

type OrderRepository interface {
	Insert(order models.Order) error
	Complete(id string) error
}

type OrderQueueRepository interface {
	Add(id, phone string) error
	Remove(id string) error
	Range(f func(k, v string) error) error
}
