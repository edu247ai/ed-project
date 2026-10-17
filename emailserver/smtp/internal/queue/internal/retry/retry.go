package retry

import (
	"context"
	"email-server/smtp/internal/queue/internal/backoff"
	"email-server/smtp/internal/queue/pkg"
)

type RetryManager struct {
	storage pkg.QueueStorage
}

func NewRetryManager(storage pkg.QueueStorage) *RetryManager {
	return &RetryManager{storage: storage}
}

func (rm *RetryManager) Fail(ctx context.Context, item pkg.QueueItem, errMsg string) error {
	next := backoff.Next(item.Attempts)
	return rm.storage.ScheduleRetry(ctx, item.ID, next, errMsg)
}

func (rm *RetryManager) Success(ctx context.Context, item pkg.QueueItem) error {
	return rm.storage.DeleteQueueItem(ctx, item.ID)
}
