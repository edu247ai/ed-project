package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var baseURL = "http://localhost:5000"

// Shared HTTP client
var client = &http.Client{
	Timeout: 15 * time.Second,
}

// -----------------------------
// Helpers
// -----------------------------

func postJSON(path string, payload interface{}) (map[string]interface{}, error) {
	body, _ := json.Marshal(payload)

	resp, err := client.Post(baseURL+path, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New("AI service error: " + resp.Status)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func getJSON(path string) (map[string]interface{}, error) {
	resp, err := client.Get(baseURL + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New("AI service error: " + resp.Status)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

// -----------------------------
// RAG
// -----------------------------

func RAGQuery(query string) (map[string]interface{}, error) {
	payload := map[string]string{"query": query}
	return postJSON("/rag/query", payload)
}

// -----------------------------
// ML: Anomaly Detection
// -----------------------------

func Anomaly(value float64) (map[string]interface{}, error) {
	payload := map[string]float64{"value": value}
	return postJSON("/ml/anomaly", payload)
}

// -----------------------------
// ML: Forecasting
// -----------------------------

func Forecast(series []float64, steps int) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"series": series,
		"steps":  steps,
	}
	return postJSON("/ml/forecast", payload)
}

// -----------------------------
// ML: Autoscale
// -----------------------------

func Autoscale(cpu float64) (map[string]interface{}, error) {
	payload := map[string]float64{"cpu": cpu}
	return postJSON("/ml/autoscale", payload)
}

// -----------------------------
// ML: Resource Optimization
// -----------------------------

func ResourceOpt(load float64, memory float64) (map[string]interface{}, error) {
	payload := map[string]float64{
		"load":   load,
		"memory": memory,
	}
	return postJSON("/ml/resource_opt", payload)
}

// -----------------------------
// RL: Evaluate Agent
// -----------------------------

func RL() (map[string]interface{}, error) {
	return getJSON("/rl/evaluate")
}

// -----------------------------
// Skills List
// -----------------------------

func Skills() ([]string, error) {
	resp, err := getJSON("/skills")
	if err != nil {
		return nil, err
	}

	skillsRaw, ok := resp["skills"].([]interface{})
	if !ok {
		return nil, errors.New("invalid skills response format")
	}

	skills := make([]string, len(skillsRaw))
	for i, v := range skillsRaw {
		s, ok := v.(string)
		if !ok {
			return nil, errors.New("invalid skill entry type")
		}
		skills[i] = s
	}

	return skills, nil
}
