package unpinned

import (
	"ipfs-server/storage/internal/blockfile"
)

type UnpinnedStore struct {
	base string
}

func New(base string) *UnpinnedStore {
	return &UnpinnedStore{base: base}
}

func (s *UnpinnedStore) Put(cid string, data []byte) error {
	return blockfile.Put(s.base, cid, data)
}

func (s *UnpinnedStore) Get(cid string) ([]byte, error) {
	return blockfile.Get(s.base, cid)
}

func (s *UnpinnedStore) Delete(cid string) error {
	return blockfile.Delete(s.base, cid)
}
