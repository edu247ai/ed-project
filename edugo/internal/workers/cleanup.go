package workers

import (
	"context"
	"log"
	"time"
)

type CleanupWorker struct{}

func NewCleanupWorker() *CleanupWorker {
	return &CleanupWorker{}
}

func (w *CleanupWorker) Name() string { return "cleanup" }

func (w *CleanupWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			log.Println("[cleanup] running cleanup tasks...")
			// TODO: delete old temp files, expired uploads, etc.
		}
	}
}
