package spam

import (
	"email-server/spam/filters/bayesian"
	"email-server/spam/filters/heuristics"
	"email-server/spam/filters/reputation"
)

type Engine struct {
	bayes *bayesian.Classifier
	rep   *reputation.Store
}

type Result struct {
	Score float64
	Flags []string
}

func NewEngine(bayes *bayesian.Classifier, rep *reputation.Store) *Engine {
	return &Engine{bayes: bayes, rep: rep}
}

func (e *Engine) Analyze(subject, body, ip, domain string) Result {
	bayesScore := e.bayes.SpamScore(body)
	heur := heuristics.Analyze(subject, body)
	repScore := e.rep.Score(ip, domain)

	total := bayesScore*5.0 + heur.Score + repScore

	flags := heur.Flags
	if bayesScore > 0.8 {
		flags = append(flags, "BAYES_HIGH")
	}
	if repScore > 3 {
		flags = append(flags, "REPUTATION_BAD")
	}

	return Result{
		Score: total,
		Flags: flags,
	}
}
