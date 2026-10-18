package unit

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"ipfs-server/pipeline"
	"ipfs-server/storage"
)

func TestPipelineUploadDownload(t *testing.T) {
	base := t.TempDir()
	st := storage.New(base)
	pl := pipeline.New(base, st)

	content := []byte("  hello\n")
	cid, err := pl.Upload(content)
	if err != nil {
		t.Fatal(err)
	}

	data, _, err := pl.Download(cid)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, content) {
		t.Fatalf("expected exact content %q, got %q", content, data)
	}
}

func TestBlockstoreValidatesIntegrityAndProtectsPinnedBlocks(t *testing.T) {
	base := t.TempDir()
	st := storage.New(base)
	content := []byte("verified block")
	pl := pipeline.New("", st)
	id, err := pl.Upload(content)
	if err != nil {
		t.Fatal(err)
	}

	if err := st.Put(id, []byte("tampered")); !errors.Is(err, storage.ErrCIDMismatch) {
		t.Fatalf("expected CID mismatch, got %v", err)
	}
	if err := st.Pin(id); err != nil {
		t.Fatal(err)
	}
	if err := st.Delete(id); !errors.Is(err, storage.ErrPinned) {
		t.Fatalf("expected pinned block protection, got %v", err)
	}
	got, err := st.Get(id)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, content) {
		t.Fatalf("expected pinned content %q, got %q", content, got)
	}

	pinnedPath := filepath.Join(base, "pinned", id)
	if _, err := os.Stat(pinnedPath); err != nil {
		t.Fatalf("pinned block was not persisted: %v", err)
	}
	if err := st.Pin(id); err != nil {
		t.Fatalf("pin should be idempotent: %v", err)
	}
	if err := st.Unpin(id); err != nil {
		t.Fatal(err)
	}
	if st.Pinned.IsPinned(id) {
		t.Fatal("expected block to be unpinned")
	}
	if err := os.WriteFile(filepath.Join(base, "unpinned", id), []byte("tampered"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := st.Get(id); !errors.Is(err, storage.ErrCIDMismatch) {
		t.Fatalf("expected read-time CID mismatch, got %v", err)
	}
}
