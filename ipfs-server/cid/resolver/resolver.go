package resolver

import (
	"errors"
	"ipfs-server/cid/validator"
	"ipfs-server/storage"
)

var ErrInvalidCID = errors.New("invalid CID")

type Resolver struct {
	store *storage.Storage
}

func New(store *storage.Storage) *Resolver {
	return &Resolver{store: store}
}

func (r *Resolver) Resolve(cid string) ([]byte, error) {
	if !validator.IsValid(cid) {
		return nil, ErrInvalidCID
	}
	return r.store.Get(cid)
}
