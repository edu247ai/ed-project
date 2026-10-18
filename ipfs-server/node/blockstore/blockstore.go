package blockstore

type Blockstore interface {
	Put(cid string, data []byte) error
	Get(cid string) ([]byte, error)
	Delete(cid string) error
}
