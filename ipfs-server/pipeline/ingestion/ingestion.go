package ingestion

import (
	"ipfs-server/cid"
	"ipfs-server/pipeline/indexing"
	"ipfs-server/storage"
)

type Ingestor struct {
	cidSvc  *cid.CIDService
	store   *storage.Storage
	indexer *indexing.Indexer
}

func New(cidSvc *cid.CIDService, store *storage.Storage, idx *indexing.Indexer) *Ingestor {
	return &Ingestor{
		cidSvc:  cidSvc,
		store:   store,
		indexer: idx,
	}
}

func (i *Ingestor) Ingest(data []byte) (string, error) {
	// 1. CID
	c := i.cidSvc.CreateFromBytes(data)

	// 2. Store block
	if err := i.store.Put(c, data); err != nil {
		return "", err
	}

	// 3. Index metadata
	i.indexer.Index(c, len(data))
	if err := i.store.Snapshots.SaveSnapshot(c, data); err != nil {
		return "", err
	}

	// (optional) pin, security, events, etc.

	return c, nil
}
