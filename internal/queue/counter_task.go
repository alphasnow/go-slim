package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

const CounterTaskName = "task:counter"

type CounterTaskPayload struct {
	Name string
}

type CounterTask struct {
	Redis *redis.Client
}

func (d *CounterTask) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload CounterTaskPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	d.Redis.Incr(ctx, "counter")
	time.Sleep(1 * time.Second)
	log.Printf("task run %v", payload)
	return nil
}
