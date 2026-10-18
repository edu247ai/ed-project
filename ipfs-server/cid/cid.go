package cid

import (
	"ipfs-server/cid/generator"
	"ipfs-server/cid/mapper"
	"ipfs-server/cid/resolver"
	"ipfs-server/cid/validator"
	"ipfs-server/storage"
)

type CIDService struct {
	Gen      *generator.Generator
	Mapper   *mapper.Mapper
	Resolver *resolver.Resolver
}

func New(base string, store *storage.Storage) *CIDService {
	gen := generator.New()
	m := mapper.New(base + "/blocks")
	res := resolver.New(store)

	return &CIDService{
		Gen:      gen,
		Mapper:   m,
		Resolver: res,
	}
}

func (c *CIDService) CreateFromBytes(data []byte) string {
	return c.Gen.FromBytes(data)
}

func (c *CIDService) IsValid(cid string) bool {
	return validator.IsValid(cid)
}
