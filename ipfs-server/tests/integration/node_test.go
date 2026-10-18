package integration

import (
	"ipfs-server/node"
	"testing"
)

func TestNodeUploadDownload(t *testing.T) {
	n := node.New("test-node", "/tmp/ipfs-node")

	cid, err := n.Upload([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}

	data, err := n.Download(cid)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "hello world" {
		t.Fatalf("expected hello world, got %s", string(data))
	}
}
