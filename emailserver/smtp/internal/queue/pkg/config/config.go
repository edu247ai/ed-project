package config

import "time"

type QueueConfig struct {
	WorkerInterval time.Duration
	MaxAttempts    int
	LocalName      string
	SMTPTargetPort string
	Timeout        time.Duration
}
