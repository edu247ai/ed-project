package pipeline

import (
	"ipfs-server/cid"
	"ipfs-server/pipeline/indexing"
	"ipfs-server/pipeline/ingestion"
	"ipfs-server/pipeline/postprocessing"
	"ipfs-server/pipeline/preprocessing"
	"ipfs-server/pipeline/retrieval"
	"ipfs-server/storage"
)

type Pipeline struct {
	Pre  *preprocessing.Preprocessor
	Ing  *ingestion.Ingestor
	Idx  *indexing.Indexer
	Post *postprocessing.Postprocessor
	Ret  *retrieval.Retriever
}

func New(base string, store *storage.Storage) *Pipeline {
	idx := indexing.New()
	cidSvc := cid.New(base, store)

	return &Pipeline{
		Pre:  preprocessing.New(),
		Ing:  ingestion.New(cidSvc, store, idx),
		Idx:  idx,
		Post: postprocessing.New(),
		Ret:  retrieval.New(cidSvc, idx),
	}
}

func (p *Pipeline) Upload(raw []byte) (string, error) {
	cleaned := p.Pre.Clean(raw)
	// you can add compression/encryption here via Post
	return p.Ing.Ingest(cleaned)
}

func (p *Pipeline) Download(cidStr string) ([]byte, indexing.Metadata, error) {
	return p.Ret.Get(cidStr)
}
