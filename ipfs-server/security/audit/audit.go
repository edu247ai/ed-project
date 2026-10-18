package audit

import (
	"encoding/json"
	"log"
	"time"
)

type AuditEvent struct {
	Timestamp time.Time              `json:"timestamp"`
	Action    string                 `json:"action"`
	CID       string                 `json:"cid,omitempty"`
	Node      string                 `json:"node,omitempty"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

type Auditor struct{}

func New() *Auditor {
	return &Auditor{}
}

func (a *Auditor) Record(event AuditEvent) {
	event.Timestamp = time.Now()
	b, _ := json.Marshal(event)
	log.Println(string(b))
}
