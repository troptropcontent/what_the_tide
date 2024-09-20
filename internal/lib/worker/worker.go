package worker

import (
	"context"

	faktory "github.com/contribsys/faktory/client"
	faktoryWorker "github.com/contribsys/faktory_worker_go"
	agenda_jobs "github.com/troptropcontent/what_the_tide/internal/modules/agenda/jobs"
)

type Job struct {
	Name     string
	Function func(ctx context.Context, args ...interface{}) error
}

func NewJob(name string, function func(ctx context.Context, args ...interface{}) error) Job {
	return Job{Name: name, Function: function}
}

func Push(name string, args ...interface{}) error {
	client, _ := faktory.Open()
	job := faktory.NewJob(name, args)
	return client.Push(job)
}

var Jobs = []Job{
	NewJob("SubscribeToAlreadyExistingAgenda", agenda_jobs.SubscribeToAlreadyExistingAgenda),
}

func NewBackGroundWorker() (manager *faktoryWorker.Manager) {
	manager = faktoryWorker.NewManager()
	for _, job := range Jobs {
		manager.Register(job.Name, job.Function)
	}
	return manager
}
