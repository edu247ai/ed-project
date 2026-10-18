package snapshots

import (
	"os"
	"path/filepath"
	"time"

	"ipfs-server/cid/validator"
)

type SnapshotStore struct {
	base string
}

func New(base string) *SnapshotStore {
	os.MkdirAll(base, 0755)
	return &SnapshotStore{base: base}
}

func (s *SnapshotStore) SaveSnapshot(cid string, data []byte) error {
	if !validator.IsValid(cid) {
		return os.ErrInvalid
	}
	if err := os.MkdirAll(s.base, 0755); err != nil {
		return err
	}
	ts := time.Now().UTC().Format("20060102T150405.000000000")
	filename := filepath.Join(s.base, cid+"-"+ts)
	return os.WriteFile(filename, data, 0644)
}
