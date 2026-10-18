package caching

import (
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	store map[string][]byte
}

func New() *Cache {
	return &Cache{
		store: make(map[string][]byte),
	}
}

func (c *Cache) Set(cid string, data []byte) {
	c.mu.Lock()
	c.store[cid] = data
	c.mu.Unlock()
}

func (c *Cache) Get(cid string) ([]byte, bool) {
	c.mu.RLock()
	v, ok := c.store[cid]
	c.mu.RUnlock()
	return v, ok
}

func (c *Cache) Delete(cid string) {
	c.mu.Lock()
	delete(c.store, cid)
	c.mu.Unlock()
}
