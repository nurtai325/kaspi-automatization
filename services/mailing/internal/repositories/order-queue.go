package repositories

import (
	"errors"
	"sync"
)

var (
	ErrNotFound               = errors.New("record is already present in the database")
	ErrIncorrectParameterType = errors.New("parameters passed to function aren't correct")
	orderQueue                = sync.Map{}
)

func OrderQueue() OrderQueueRepository {
	return &orderQueueRepository{}
}

type orderQueueRepository struct {
}

func (o *orderQueueRepository) Add(id, phone string) error {
	orderQueue.Store(id, phone)
	return nil
}

func (o *orderQueueRepository) Remove(id string) error {
	orderQueue.Delete(id)
	return nil
}

func (o *orderQueueRepository) Range(f func(k, v string) error) error {
	var rangeErr error

	orderQueue.Range(func(k, v any) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}

		value, ok := k.(string)
		if !ok {
			rangeErr = ErrIncorrectParameterType
			return false
		}

		err := f(key, value)
		if err != nil {
			rangeErr = err
			return false
		}

		return true
	})

	return rangeErr
}
