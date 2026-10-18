package gc

import (
	"os"
	"path/filepath"
)

type GC struct {
	unpinnedPath string
	pinnedPath   string
}

func New(unpinned, pinned string) *GC {
	return &GC{unpinnedPath: unpinned, pinnedPath: pinned}
}

func (g *GC) Run() error {
	entries, err := os.ReadDir(g.unpinnedPath)
	if err != nil {
		return err
	}

	for _, e := range entries {
		cid := e.Name()
		pinned := filepath.Join(g.pinnedPath, cid)

		if _, err := os.Stat(pinned); os.IsNotExist(err) {
			os.Remove(filepath.Join(g.unpinnedPath, cid))
		}
	}

	return nil
}
