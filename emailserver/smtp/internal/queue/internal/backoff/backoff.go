package backoff

import "time"

var schedule = []time.Duration{
	time.Minute,
	5 * time.Minute,
	15 * time.Minute,
	time.Hour,
	4 * time.Hour,
	12 * time.Hour,
}

func Next(attempts int) time.Time {
	if attempts >= len(schedule) {
		attempts = len(schedule) - 1
	}
	return time.Now().Add(schedule[attempts])
}
