package workers

import (
	"context"
	"email-server/smtp/internal/queue/internal/retry"
	"email-server/smtp/internal/queue/pkg"
	"email-server/smtp/pkg/logger"
	"time"
)

type Worker struct {
	storage    pkg.QueueStorage
	dispatcher *Dispatcher
	retryMgr   *retry.RetryManager
	interval   time.Duration
}

func NewWorker(storage pkg.QueueStorage, dispatcher *Dispatcher, retryMgr *retry.RetryManager, interval time.Duration) *Worker {
	return &Worker{
		storage:    storage,
		dispatcher: dispatcher,
		retryMgr:   retryMgr,
		interval:   interval,
	}
}

func (w *Worker) Run(ctx context.Context) {
	if w.interval <= 0 {
		w.interval = time.Minute
	}
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	w.process(ctx)
	for {
		select {
		case <-ticker.C:
			w.process(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (w *Worker) process(ctx context.Context) {
	items, err := w.storage.LoadQueueItems(ctx)
	if err != nil {
		logger.Error("queue load error: %v", err)
		return
	}

	now := time.Now()

	for _, item := range items {
		if item.NextAttempt.After(now) {
			continue
		}

		if err := w.dispatcher.Deliver(ctx, item); err != nil {
			logger.Error("delivery error: %v", err)
			if retryErr := w.retryMgr.Fail(ctx, item, err.Error()); retryErr != nil {
				logger.Error("queue retry scheduling failed: %v", retryErr)
			}
			continue
		}
		if err := w.retryMgr.Success(ctx, item); err != nil {
			logger.Error("queue completion failed: %v", err)
		}
	}
}
