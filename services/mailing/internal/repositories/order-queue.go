package repositories

import (
	"errors"
	"sync"

	"github.com/nurtai325/kaspi/mailing/internal/models"
)

var (
	ErrNotFound               = errors.New("record is already present in the database")
	ErrIncorrectParameterType = errors.New("parameters passed to function aren't correct")
	ErrOrderQueueValue        = errors.New(`value in the orders queue is not of format "productCode_token"`)
	orderQueue                = sync.Map{}
)

func OrderQueue() OrderQueueRepository {
	return &orderQueueRepository{}
}

type orderQueueRepository struct {
}

func (o *orderQueueRepository) Add(id string, order models.QueuedOrder) error {
	orderQueue.Store(id, order)
	return nil
}

func (o *orderQueueRepository) Remove(id string) error {
	orderQueue.Delete(id)
	return nil
}

func (o *orderQueueRepository) Range(f func(id string, order models.QueuedOrder) error) error {
	var rangeErr error

	orderQueue.Range(func(k, v any) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}

		value, ok := v.(models.QueuedOrder)
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
