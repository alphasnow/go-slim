package queue

import (
	"github.com/hibiken/asynq"
	"go-slim/pkg/xredis"
)

func GetRedisClientOpt(rds *xredis.Config) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:     rds.GetAddr(),
		Password: rds.Pass,
		DB:       rds.DB,
	}
}

func GetConfig(logger *Logger) asynq.Config {
	return asynq.Config{
		// 同时运行的任务数量,默认cpu数量
		// Concurrency: 6,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{
			QueueCritical: 6,
			QueueDefault:  3,
			QueueLow:      1,
		},
		StrictPriority: true,
		Logger:         logger,
	}
}
