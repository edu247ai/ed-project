package peers

type Peer struct {
	ID      string
	Address string
}

type PeerStore struct {
	peers map[string]Peer
}

func New() *PeerStore {
	return &PeerStore{peers: make(map[string]Peer)}
}

func (ps *PeerStore) Add(p Peer) {
	ps.peers[p.ID] = p
}

func (ps *PeerStore) List() []Peer {
	out := []Peer{}
	for _, p := range ps.peers {
		out = append(out, p)
	}
	return out
}
