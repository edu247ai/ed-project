package pinned

import (
	"ipfs-server/storage/internal/blockfile"
)

type PinnedStore struct {
	base string
}

func New(base string) *PinnedStore {
	return &PinnedStore{base: base}
}

func (s *PinnedStore) Pin(cid string, data []byte) error {
	return blockfile.Put(s.base, cid, data)
}

func (s *PinnedStore) Get(cid string) ([]byte, error) {
	return blockfile.Get(s.base, cid)
}

func (s *PinnedStore) Unpin(cid string) error {
	return blockfile.Delete(s.base, cid)
}

func (s *PinnedStore) IsPinned(cid string) bool {
	return blockfile.Exists(s.base, cid)
}
