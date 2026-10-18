package datastore

import "sync"

type Datastore struct {
	mu   sync.RWMutex
	meta map[string]map[string]interface{}
}

func New() *Datastore {
	return &Datastore{
		meta: make(map[string]map[string]interface{}),
	}
}

func (ds *Datastore) Set(cid string, key string, value interface{}) {
	ds.mu.Lock()
	if _, ok := ds.meta[cid]; !ok {
		ds.meta[cid] = make(map[string]interface{})
	}
	ds.meta[cid][key] = value
	ds.mu.Unlock()
}

func (ds *Datastore) Get(cid string) map[string]interface{} {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.meta[cid]
}
