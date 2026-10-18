package cluster

import "sync"

type NodeInfo struct {
	ID      string
	Address string
	Healthy bool
}

type Cluster struct {
	mu    sync.RWMutex
	nodes map[string]NodeInfo
}

func New() *Cluster {
	return &Cluster{
		nodes: make(map[string]NodeInfo),
	}
}

func (c *Cluster) Register(id, addr string) {
	c.mu.Lock()
	c.nodes[id] = NodeInfo{ID: id, Address: addr, Healthy: true}
	c.mu.Unlock()
}

func (c *Cluster) MarkUnhealthy(id string) {
	c.mu.Lock()
	n := c.nodes[id]
	n.Healthy = false
	c.nodes[id] = n
	c.mu.Unlock()
}

func (c *Cluster) List() []NodeInfo {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := []NodeInfo{}
	for _, n := range c.nodes {
		out = append(out, n)
	}
	return out
}
