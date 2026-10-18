package unit

import (
	"ipfs-server/distribution/sharding"
	"testing"
)

func TestSharding(t *testing.T) {
	nodes := []string{"node1", "node2", "node3"}
	s := sharding.New(nodes)

	cid := "abc123"
	node := s.PickNode(cid)

	if node == "" {
		t.Fatal("expected a node, got empty")
	}
}
