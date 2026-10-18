package node

import (
	"log"

	"ipfs-server/config"
	"ipfs-server/monitoring"
	"ipfs-server/pipeline"
	"ipfs-server/security"
	"ipfs-server/storage"
)

type Node struct {
	ID       string
	BasePath string

	Storage  *storage.Storage
	Pipeline *pipeline.Pipeline
	Security *security.Security
	Monitor  *monitoring.Monitoring
}

func New(id string, base string) *Node {
	st := storage.New(base + "/storage")
	pl := pipeline.New(base, st)

	return &Node{
		ID:       id,
		BasePath: base,
		Storage:  st,
		Pipeline: pl,
	}
}

func (n *Node) Upload(data []byte) (string, error) {
	cid, err := n.Pipeline.Upload(data)
	if err != nil {
		return "", err
	}
	log.Printf("node=%s upload cid=%s\n", n.ID, cid)
	return cid, nil
}

func (n *Node) Download(cidStr string) ([]byte, error) {
	data, _, err := n.Pipeline.Download(cidStr)
	if err != nil {
		return nil, err
	}
	log.Printf("node=%s download cid=%s\n", n.ID, cidStr)
	return data, nil
}

func NewWithConfig(id string, cfg *config.Config, sec *security.Security, mon *monitoring.Monitoring) *Node {
	st := storage.New(cfg.Storage.Base)
	pl := pipeline.New(cfg.Storage.Base, st)

	return &Node{
		ID:       id,
		BasePath: cfg.Storage.Base,
		Storage:  st,
		Pipeline: pl,
		Security: sec,
		Monitor:  mon,
	}
}
