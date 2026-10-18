package sharding

import "hash/fnv"

type Sharder struct {
	nodes []string // node IDs or addresses
}

func New(nodes []string) *Sharder {
	return &Sharder{nodes: nodes}
}

func (s *Sharder) PickNode(cid string) string {
	h := fnv.New32a()
	h.Write([]byte(cid))
	idx := int(h.Sum32()) % len(s.nodes)
	return s.nodes[idx]
}
