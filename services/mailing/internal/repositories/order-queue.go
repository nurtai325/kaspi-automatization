package repositories

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("record is already present in the database")
	orderQueue  = sync.Map{}
)

func OrderQueue() OrderQueueRepository {
	return orderQueueRepository{}
}

type orderQueueRepository struct {
}

func (o orderQueueRepository) Add(id, phone string) error {
	orderQueue.Store(id, phone)
	return nil
}
