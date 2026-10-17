package heuristics

import "strings"

type Result struct {
	Score float64
	Flags []string
}

func Analyze(subject, body string) Result {
	score := 0.0
	flags := []string{}

	if strings.Contains(strings.ToLower(subject), "free") {
		score += 1.5
		flags = append(flags, "SUBJECT_FREE")
	}
	if strings.Contains(strings.ToLower(body), "click here") {
		score += 1.0
		flags = append(flags, "BODY_CLICK_HERE")
	}
	if strings.Contains(strings.ToLower(body), "winner") {
		score += 1.0
		flags = append(flags, "BODY_WINNER")
	}

	return Result{Score: score, Flags: flags}
}
