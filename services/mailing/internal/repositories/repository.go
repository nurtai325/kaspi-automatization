package repositories

import "github.com/nurtai325/kaspi/mailing/internal/models"

type OrderRepository interface {
	GetById(id string) (models.Order, error)
	Exists(id string) (bool, error)
	Insert(order models.Order) error
	Complete(id string) error
}
