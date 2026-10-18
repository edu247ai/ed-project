package swarm

import (
	"io"
	"net/http"
)

type Swarm struct{}

func New() *Swarm {
	return &Swarm{}
}

func (s *Swarm) FetchBlock(peerAddr, cid string) ([]byte, error) {
	resp, err := http.Get(peerAddr + "/block/" + cid)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
