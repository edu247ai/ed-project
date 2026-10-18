package controllers

import (
	"encoding/json"
	"net/http"

	"ipfs-server/api/dto"
	"ipfs-server/node"
)

type UploadController struct {
	Node *node.Node
}

func NewUploadController(n *node.Node) *UploadController {
	return &UploadController{Node: n}
}

func (c *UploadController) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.UploadRequest
	json.NewDecoder(r.Body).Decode(&req)

	cid, err := c.Node.Upload(req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"cid": cid})
}
