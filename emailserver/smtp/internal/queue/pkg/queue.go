package pkg

import (
	"context"
	"time"
)

type QueueItem struct {
	ID          string
	MessageID   string
	Recipient   string
	Attempts    int
	NextAttempt time.Time
}

type QueueStorage interface {
	LoadQueueItems(context.Context) ([]QueueItem, error)
	LoadMessage(context.Context, string) ([]byte, string, error)
	ScheduleRetry(context.Context, string, time.Time, string) error
	DeleteQueueItem(context.Context, string) error
}
