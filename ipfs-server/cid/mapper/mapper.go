package mapper

import "path/filepath"

type Mapper struct {
	basePath string
}

func New(base string) *Mapper {
	return &Mapper{basePath: base}
}

// Local path mapping (can later be extended to shard/node mapping)
func (m *Mapper) Path(cid string) string {
	return filepath.Join(m.basePath, cid)
}
