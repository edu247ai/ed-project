package libp2p

type Transport struct {
	NodeID string
}

func New(nodeID string) *Transport {
	return &Transport{NodeID: nodeID}
}

func (t *Transport) Broadcast(msg []byte) error {
	// placeholder for real libp2p broadcast
	return nil
}
