package repositories

import "errors"

var (
	ErrNotFound = errors.New("record is already present in the database")
)

func OrderQueue() OrderQueueRepository {
	return nil
}
