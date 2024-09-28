package tasks

import (
	scheduling "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

type stopTasks func()

func Start(conf models.Config) (stopTasks, error) {
	scheduler := scheduling.New()

	tasks := []*scheduling.Task{orderTask(conf.KASPI_TOKEN)}
	for _, task := range tasks {
		_, err := scheduler.Add(task)
		if err != nil {
			return scheduler.Stop, err
		}
	}

	return scheduler.Stop, nil
}

func errFunc(err error) {
	panic(err)
}
