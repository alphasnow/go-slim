package queue

import "github.com/hibiken/asynq"

type Tasks struct {
	CounterTask *CounterTask
}

func (ts *Tasks) GetHandlerMap() map[string]asynq.Handler {
	return map[string]asynq.Handler{
		CounterTaskName: ts.CounterTask,
	}
}
