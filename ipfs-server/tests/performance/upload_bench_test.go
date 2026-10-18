package performance

import (
	"ipfs-server/node"
	"testing"
)

func BenchmarkUpload(b *testing.B) {
	n := node.New("bench-node", "/tmp/ipfs-bench")

	data := []byte("hello world")

	for i := 0; i < b.N; i++ {
		_, _ = n.Upload(data)
	}
}
