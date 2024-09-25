package worker

import (
	"context"

	faktory "github.com/contribsys/faktory/client"
	faktoryWorker "github.com/contribsys/faktory_worker_go"
	calendar_jobs "github.com/troptropcontent/what_the_tide/internal/modules/calendar/jobs"
)

type Job interface {
	Name() string
	Perform(ctx context.Context, args ...interface{}) error
}

func Push(j Job, args ...interface{}) error {
	client, _ := faktory.Open()
	job := faktory.NewJob(j.Name(), args)
	return client.Push(job)
}

var Jobs = []Job{
	&calendar_jobs.SubscribeToAlreadyExistingAgendaJob{},
}

func NewBackGroundWorker() (manager *faktoryWorker.Manager) {
	manager = faktoryWorker.NewManager()
	for _, job := range Jobs {
		manager.Register(job.Name(), job.Perform)
	}
	return manager
}
