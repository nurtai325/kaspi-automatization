package tasks

import (
	scheduling "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

type stopTasks func()

func Start(conf models.Config, tasks []*scheduling.Task) (stopTasks, error) {
	scheduler := scheduling.New()

	for _, task := range tasks {
		_, err := scheduler.Add(task)
		if err != nil {
			return scheduler.Stop, err
		}
	}

	return scheduler.Stop, nil
}
