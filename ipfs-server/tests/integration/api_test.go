package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ipfs-server/api/dto"
	"ipfs-server/api/routes"
	"ipfs-server/node"
)

func TestAPIUploadDownload(t *testing.T) {
	n := node.New("test-node", t.TempDir())
	server := httptest.NewServer(routes.NewHandler(n))
	defer server.Close()

	// Upload
	body := `{"data":"aGVsbG8="}`
	req, _ := http.NewRequest("POST", server.URL+"/upload", bytes.NewBuffer([]byte(body)))
	req.Header.Set("X-API-Key", "super-secret-key")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("upload returned status %d", resp.StatusCode)
	}
	var out map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}

	cid := out["cid"]
	if cid == "" {
		t.Fatal("upload response did not include a CID")
	}

	// Download
	req2, _ := http.NewRequest("GET", server.URL+"/download?cid="+cid, nil)
	req2.Header.Set("X-API-Key", "super-secret-key")

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		t.Fatal(err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusOK {
		t.Fatalf("download returned status %d", resp2.StatusCode)
	}
	var out2 dto.DownloadResponse
	if err := json.NewDecoder(resp2.Body).Decode(&out2); err != nil {
		t.Fatal(err)
	}

	if out2.CID != cid {
		t.Fatalf("expected cid %s, got %s", cid, out2.CID)
	}
	if out2.Size != len("hello") || !bytes.Equal(out2.Data, []byte("hello")) {
		t.Fatalf("download returned size %d and content %q, want %q", out2.Size, out2.Data, "hello")
	}
}
