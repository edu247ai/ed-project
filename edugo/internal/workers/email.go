package workers

import (
	"context"
	"log"
	"time"
)

type EmailWorker struct{}

func NewEmailWorker() *EmailWorker {
	return &EmailWorker{}
}

func (w *EmailWorker) Name() string { return "email" }

func (w *EmailWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			log.Println("[email] sending queued emails...")
			// TODO: send queued emails from DB
		}
	}
}
