package repositories

import "github.com/nurtai325/kaspi/mailing/internal/models"

type OrderRepository interface {
	Insert(order models.Order) error
	Complete(id string) error
	Exists(id string) (bool, error)
}

type OrderQueueRepository interface {
	Add(id string, order models.QueuedOrder) error
	Remove(id string) error
	Range(f func(id string, order models.QueuedOrder) error) error
}

type ClientRepository interface {
	Get() ([]models.Client, error)
	Insert(models.Client) error
	Extend(id int, duration int, unit string) error
	ConnectWh(id int) error
	Deactivate(id int) error
}
