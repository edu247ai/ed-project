package workers

import (
	"context"
	"log"
	"sync"
)

type Worker interface {
	Name() string
	Run(ctx context.Context)
}

type Supervisor struct {
	workers []Worker
}

func NewSupervisor() *Supervisor {
	return &Supervisor{
		workers: []Worker{},
	}
}

func (s *Supervisor) Register(w Worker) {
	s.workers = append(s.workers, w)
}

func (s *Supervisor) Start(ctx context.Context) {
	var wg sync.WaitGroup

	for _, w := range s.workers {
		wg.Add(1)
		go func(w Worker) {
			defer wg.Done()
			log.Printf("[worker:%s] started", w.Name())
			w.Run(ctx)
			log.Printf("[worker:%s] stopped", w.Name())
		}(w)
	}

	go func() {
		<-ctx.Done()
		log.Println("Supervisor received shutdown signal")
	}()

	wg.Wait()
}
