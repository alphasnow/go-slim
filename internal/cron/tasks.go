package cron

import (
	"github.com/go-co-op/gocron"
)

type Tasks struct {
	CleanTmpFileTask *CleanTmpFileTask
}

func (ts *Tasks) Handle(scheduler *gocron.Scheduler) {

	scheduler.Every(30).Minutes().Do(ts.CleanTmpFileTask.Handle)
}
