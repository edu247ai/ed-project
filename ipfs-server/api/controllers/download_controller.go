package controllers

import (
	"encoding/json"
	"net/http"

	"ipfs-server/api/dto"
	"ipfs-server/node"
)

type DownloadController struct {
	Node *node.Node
}

func NewDownloadController(n *node.Node) *DownloadController {
	return &DownloadController{Node: n}
}

func (c *DownloadController) Handle(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	data, err := c.Node.Download(cid)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	resp := dto.DownloadResponse{
		CID:  cid,
		Size: len(data),
		Data: data,
	}

	json.NewEncoder(w).Encode(resp)
}
