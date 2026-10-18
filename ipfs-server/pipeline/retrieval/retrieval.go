package retrieval

import (
	"ipfs-server/cid"
	"ipfs-server/pipeline/indexing"
)

type Retriever struct {
	cidSvc  *cid.CIDService
	indexer *indexing.Indexer
}

func New(cidSvc *cid.CIDService, idx *indexing.Indexer) *Retriever {
	return &Retriever{
		cidSvc:  cidSvc,
		indexer: idx,
	}
}

func (r *Retriever) Get(cidStr string) ([]byte, indexing.Metadata, error) {
	data, err := r.cidSvc.Resolver.Resolve(cidStr)
	if err != nil {
		return nil, indexing.Metadata{}, err
	}

	meta, _ := r.indexer.Get(cidStr)
	return data, meta, nil
}
