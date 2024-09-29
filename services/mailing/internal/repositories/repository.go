package repositories

import "github.com/nurtai325/kaspi/mailing/internal/models"

type Order interface {
	GetById(id string) (models.Order, error)
	Inser(order models.Order) error
	Complete(id string) error
}
