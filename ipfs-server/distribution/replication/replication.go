package replication

import (
	"bytes"
	"net/http"
)

type Replicator struct {
	nodeEndpoints []string // e.g. http://node2:8080
}

func New(endpoints []string) *Replicator {
	return &Replicator{nodeEndpoints: endpoints}
}

func (r *Replicator) Replicate(cid string, data []byte) {
	for _, ep := range r.nodeEndpoints {
		// fire-and-forget replication
		_, _ = http.Post(ep+"/replicate/"+cid, "application/octet-stream", bytes.NewReader(data))
	}
}
