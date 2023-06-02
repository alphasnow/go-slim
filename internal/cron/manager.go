package cron

import (
	"github.com/go-co-op/gocron"
	"time"
)

type Manager struct {
	Tasks     *Tasks
	scheduler *gocron.Scheduler
}

func (m *Manager) newScheduler() *gocron.Scheduler {
	// https://github.com/go-co-op/gocron
	scheduler := gocron.NewScheduler(time.Local)
	// CronScheduler.SetMaxConcurrentJobs()

	m.Tasks.Handle(scheduler)

	return scheduler
}

func (m *Manager) StartAsync() *gocron.Scheduler {

	scheduler := m.newScheduler()
	scheduler.StartAsync()

	m.scheduler = scheduler
	return m.scheduler
}

func (m *Manager) Close() {
	m.scheduler.Stop()
}
