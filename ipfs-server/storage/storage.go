package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"ipfs-server/cid/generator"
	"ipfs-server/cid/validator"
	"ipfs-server/node/blockstore"
	"ipfs-server/storage/gc"
	"ipfs-server/storage/pinned"
	"ipfs-server/storage/snapshots"
	"ipfs-server/storage/temp"
	"ipfs-server/storage/unpinned"
)

type Storage struct {
	mu        sync.RWMutex
	Unpinned  *unpinned.UnpinnedStore
	Pinned    *pinned.PinnedStore
	Temp      *temp.TempStore
	Snapshots *snapshots.SnapshotStore
	GC        *gc.GC
}

var _ blockstore.Blockstore = (*Storage)(nil)

var (
	ErrInvalidCID  = errors.New("invalid CID")
	ErrCIDMismatch = errors.New("block does not match CID")
	ErrPinned      = errors.New("block is pinned")
)

func New(base string) *Storage {
	return &Storage{
		Unpinned:  unpinned.New(filepath.Join(base, "unpinned")),
		Pinned:    pinned.New(filepath.Join(base, "pinned")),
		Temp:      temp.New(filepath.Join(base, "temp")),
		Snapshots: snapshots.New(filepath.Join(base, "snapshots")),
		GC:        gc.New(filepath.Join(base, "unpinned"), filepath.Join(base, "pinned")),
	}
}

func (s *Storage) Put(cid string, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !validator.IsValid(cid) {
		return ErrInvalidCID
	}
	if generator.New().FromBytes(data) != cid {
		return ErrCIDMismatch
	}
	if s.Pinned.IsPinned(cid) {
		pinnedData, err := s.Pinned.Get(cid)
		if err != nil {
			return err
		}
		if generator.New().FromBytes(pinnedData) != cid {
			return fmt.Errorf("%w: %s", ErrCIDMismatch, cid)
		}
		return nil
	}
	return s.Unpinned.Put(cid, data)
}

func (s *Storage) Get(cid string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.get(cid)
}

func (s *Storage) get(cid string) ([]byte, error) {
	if !validator.IsValid(cid) {
		return nil, ErrInvalidCID
	}

	var (
		data []byte
		err  error
	)
	if s.Pinned.IsPinned(cid) {
		data, err = s.Pinned.Get(cid)
	} else {
		data, err = s.Unpinned.Get(cid)
	}
	if err != nil {
		return nil, err
	}
	if generator.New().FromBytes(data) != cid {
		return nil, fmt.Errorf("%w: %s", ErrCIDMismatch, cid)
	}
	return data, nil
}

func (s *Storage) Delete(cid string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !validator.IsValid(cid) {
		return ErrInvalidCID
	}
	if s.Pinned.IsPinned(cid) {
		return ErrPinned
	}
	return s.Unpinned.Delete(cid)
}

func (s *Storage) Pin(cid string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !validator.IsValid(cid) {
		return ErrInvalidCID
	}
	data, err := s.get(cid)
	if err != nil {
		return err
	}
	if s.Pinned.IsPinned(cid) {
		return nil
	}
	if err := s.Pinned.Pin(cid, data); err != nil {
		return err
	}
	if err := s.Unpinned.Delete(cid); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func (s *Storage) Unpin(cid string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !validator.IsValid(cid) {
		return ErrInvalidCID
	}
	data, err := s.Pinned.Get(cid)
	if err != nil {
		return err
	}
	if generator.New().FromBytes(data) != cid {
		return fmt.Errorf("%w: %s", ErrCIDMismatch, cid)
	}
	if err := s.Unpinned.Put(cid, data); err != nil {
		return err
	}
	return s.Pinned.Unpin(cid)
}
