package spam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RspamdClient struct {
	endpoint string
	client   *http.Client
}

func NewRspamdClient(endpoint string) *RspamdClient {
	return &RspamdClient{
		endpoint: endpoint,
		client:   &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *RspamdClient) Check(ctx context.Context, raw []byte) (float64, []string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(raw))
	if err != nil {
		return 0, nil, fmt.Errorf("create Rspamd request: %w", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("call Rspamd checkv2: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4096))
		return 0, nil, fmt.Errorf("Rspamd checkv2 returned HTTP %s", resp.Status)
	}
	var result struct {
		Score   float64            `json:"score"`
		Symbols map[string]float64 `json:"symbols"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&result); err != nil {
		return 0, nil, fmt.Errorf("decode Rspamd response: %w", err)
	}
	flags := make([]string, 0, len(result.Symbols))
	for symbol := range result.Symbols {
		flags = append(flags, symbol)
	}
	return result.Score, flags, nil
}

func CheckRspamd(raw []byte) (float64, []string, error) {
	return NewRspamdClient("http://127.0.0.1:11333/checkv2").Check(context.Background(), raw)
}
