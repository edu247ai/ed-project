package indexing

import (
	"sync"
	"time"
)

type Metadata struct {
	CID       string
	Size      int
	Timestamp time.Time
}

type Indexer struct {
	mu sync.RWMutex
	db map[string]Metadata
}

func New() *Indexer {
	return &Indexer{db: make(map[string]Metadata)}
}

func (idx *Indexer) Index(cid string, size int) {
	idx.mu.Lock()
	idx.db[cid] = Metadata{
		CID:       cid,
		Size:      size,
		Timestamp: time.Now(),
	}
	idx.mu.Unlock()
}

func (idx *Indexer) Get(cid string) (Metadata, bool) {
	idx.mu.RLock()
	m, ok := idx.db[cid]
	idx.mu.RUnlock()
	return m, ok
}
