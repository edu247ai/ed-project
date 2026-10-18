package daemon

import (
	"ipfs-server/node"
	"log"
)

type Daemon struct {
	node *node.Node
}

func New(n *node.Node) *Daemon {
	return &Daemon{node: n}
}

func (d *Daemon) Start() {
	log.Println("Node daemon started:", d.node.ID)
}

func (d *Daemon) Node() *node.Node {
	return d.node
}
