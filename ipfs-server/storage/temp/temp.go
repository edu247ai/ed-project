package temp

import (
	"os"
	"path/filepath"
)

type TempStore struct {
	base string
}

func New(base string) *TempStore {
	os.MkdirAll(base, 0755)
	return &TempStore{base: base}
}

func (t *TempStore) SaveTemp(name string, data []byte) error {
	return os.WriteFile(filepath.Join(t.base, name), data, 0644)
}

func (t *TempStore) LoadTemp(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(t.base, name))
}

func (t *TempStore) DeleteTemp(name string) error {
	return os.Remove(filepath.Join(t.base, name))
}
