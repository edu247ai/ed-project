package metrics

import (
	"sync"
	"time"
)

type Counter struct {
	mu   sync.Mutex
	name string
	val  int64
}

func NewCounter(name string) *Counter {
	return &Counter{name: name}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}

func (c *Counter) Add(n int64) {
	c.mu.Lock()
	c.val += n
	c.mu.Unlock()
}

func (c *Counter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

type Timer struct {
	mu    sync.Mutex
	name  string
	times []time.Duration
}

func NewTimer(name string) *Timer {
	return &Timer{name: name}
}

func (t *Timer) Observe(d time.Duration) {
	t.mu.Lock()
	t.times = append(t.times, d)
	t.mu.Unlock()
}

func (t *Timer) Avg() time.Duration {
	t.mu.Lock()
	defer t.mu.Unlock()
	var total time.Duration
	for _, v := range t.times {
		total += v
	}
	if len(t.times) == 0 {
		return 0
	}
	return total / time.Duration(len(t.times))
}
