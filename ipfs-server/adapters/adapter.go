package adapters

import (
	"ipfs-server/adapters/cluster"
	"ipfs-server/adapters/db"
	"ipfs-server/adapters/external"
	"ipfs-server/adapters/redis"
	"ipfs-server/adapters/s3"
)

type Adapters struct {
	Cluster  *cluster.Cluster
	DB       *db.DB
	External func(string) ([]byte, error)
	Redis    *redis.Client
	S3       *s3.S3Adapter
}

func New(cluster *cluster.Cluster, db *db.DB, redis *redis.Client, s3 *s3.S3Adapter) *Adapters {
	return &Adapters{
		Cluster:  cluster,
		DB:       db,
		External: external.FetchURL,
		Redis:    redis,
		S3:       s3,
	}
}
