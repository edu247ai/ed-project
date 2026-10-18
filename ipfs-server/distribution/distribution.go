package distribution

import (
	"ipfs-server/distribution/caching"
	loadbalancing "ipfs-server/distribution/load-balancing"
	"ipfs-server/distribution/replication"
	"ipfs-server/distribution/sharding"
)

type Distribution struct {
	Cache      *caching.Cache
	Balancer   *loadbalancing.Balancer
	Sharder    *sharding.Sharder
	Replicator *replication.Replicator
}

func New(nodes []loadbalancing.NodeInfo, shardNodes []string, replicaEndpoints []string) *Distribution {
	return &Distribution{
		Cache:      caching.New(),
		Balancer:   loadbalancing.New(nodes),
		Sharder:    sharding.New(shardNodes),
		Replicator: replication.New(replicaEndpoints),
	}
}
