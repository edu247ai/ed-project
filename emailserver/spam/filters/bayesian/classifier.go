package bayesian

import (
	"math"
	"strings"
)

type Classifier struct {
	hamCounts  map[string]int
	spamCounts map[string]int
	hamTotal   int
	spamTotal  int
}

func NewClassifier() *Classifier {
	return &Classifier{
		hamCounts:  make(map[string]int),
		spamCounts: make(map[string]int),
	}
}

func tokenize(body string) []string {
	body = strings.ToLower(body)
	return strings.Fields(body)
}

func (c *Classifier) TrainHam(body string) {
	for _, t := range tokenize(body) {
		c.hamCounts[t]++
		c.hamTotal++
	}
}

func (c *Classifier) TrainSpam(body string) {
	for _, t := range tokenize(body) {
		c.spamCounts[t]++
		c.spamTotal++
	}
}

func (c *Classifier) SpamScore(body string) float64 {
	tokens := tokenize(body)

	// naive Bayes log-likelihood
	var logHam, logSpam float64

	for _, t := range tokens {
		ham := float64(c.hamCounts[t]+1) / float64(c.hamTotal+len(c.hamCounts))
		spam := float64(c.spamCounts[t]+1) / float64(c.spamTotal+len(c.spamCounts))

		logHam += math.Log(ham)
		logSpam += math.Log(spam)
	}

	// convert to probability
	denom := math.Exp(logHam) + math.Exp(logSpam)
	if denom == 0 {
		return 0.5
	}
	return math.Exp(logSpam) / denom
}
