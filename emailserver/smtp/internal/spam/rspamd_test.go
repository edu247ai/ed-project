package spam

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRspamdClientUsesCheckv2HTTPResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/checkv2" {
			t.Errorf("request = %s %s, want POST /checkv2", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"score":9.2,"symbols":{"SPAM":2.1,"PHISHING":4.0}}`)
	}))
	defer server.Close()

	score, symbols, err := NewRspamdClient(server.URL+"/checkv2").Check(context.Background(), []byte("message"))
	if err != nil {
		t.Fatal(err)
	}
	if score != 9.2 || len(symbols) != 2 {
		t.Fatalf("Check() = (%v, %v), want score 9.2 and 2 symbols", score, symbols)
	}
}
