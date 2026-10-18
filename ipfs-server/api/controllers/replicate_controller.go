package controllers

import (
	"io"
	"net/http"

	"ipfs-server/node"
)

type ReplicateController struct {
	Node *node.Node
}

func NewReplicateController(n *node.Node) *ReplicateController {
	return &ReplicateController{Node: n}
}

func (c *ReplicateController) Handle(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	data, _ := io.ReadAll(r.Body)

	// Store replicated block
	c.Node.Storage.Unpinned.Put(cid, data)

	w.WriteHeader(http.StatusOK)
}
