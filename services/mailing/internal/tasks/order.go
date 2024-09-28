package tasks

import (
	"time"

	schedule "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/external/kaspi"
)

func orderTask(token string) *schedule.Task {
	return &schedule.Task{
		Interval: 3 * time.Minute,
		TaskFunc: func() error {
			return kaspi.RefhreshOrders(token)
		},
		ErrFunc: errFunc,
	}
}
