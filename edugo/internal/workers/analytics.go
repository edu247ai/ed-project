package workers

import (
	"context"
	"log"
	"time"
)

type AnalyticsWorker struct{}

func NewAnalyticsWorker() *AnalyticsWorker {
	return &AnalyticsWorker{}
}

func (w *AnalyticsWorker) Name() string { return "analytics" }

func (w *AnalyticsWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			log.Println("[analytics] computing metrics...")
			// TODO: compute course stats, user activity, etc.
		}
	}
}
