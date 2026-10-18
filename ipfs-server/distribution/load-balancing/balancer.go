package loadbalancing

import "math/rand"

type NodeInfo struct {
	ID      string
	Address string
	Load    int
}

type Balancer struct {
	nodes []NodeInfo
}

func New(nodes []NodeInfo) *Balancer {
	return &Balancer{nodes: nodes}
}

// Random strategy
func (b *Balancer) PickRandom() NodeInfo {
	return b.nodes[rand.Intn(len(b.nodes))]
}

// Least-loaded strategy
func (b *Balancer) PickLeastLoaded() NodeInfo {
	best := b.nodes[0]
	for _, n := range b.nodes {
		if n.Load < best.Load {
			best = n
		}
	}
	return best
}
