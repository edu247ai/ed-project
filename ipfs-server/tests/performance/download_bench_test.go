package performance

import (
	"ipfs-server/node"
	"testing"
)

func BenchmarkDownload(b *testing.B) {
	n := node.New("bench-node", "/tmp/ipfs-bench")

	cid, _ := n.Upload([]byte("hello world"))

	for i := 0; i < b.N; i++ {
		_, _ = n.Download(cid)
	}
}
