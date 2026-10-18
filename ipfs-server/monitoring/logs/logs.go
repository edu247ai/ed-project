package logs

import (
	"encoding/json"
	"log"
)

func Info(event string, fields map[string]interface{}) {
	entry := map[string]interface{}{
		"level":  "INFO",
		"event":  event,
		"fields": fields,
	}
	b, _ := json.Marshal(entry)
	log.Println(string(b))
}

func Error(event string, err error, fields map[string]interface{}) {
	entry := map[string]interface{}{
		"level":  "ERROR",
		"event":  event,
		"error":  err.Error(),
		"fields": fields,
	}
	b, _ := json.Marshal(entry)
	log.Println(string(b))
}
