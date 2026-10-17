package tests

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"email-server/smtp/internal/handlers"
	"email-server/smtp/internal/server"
	"email-server/smtp/internal/storage"
	"email-server/smtp/pkg/config"
)

type spamChecker float64

func (s spamChecker) Check(context.Context, []byte) (float64, []string, error) {
	return float64(s), nil, nil
}

func TestSMTPAcceptsAndQueuesMessage(t *testing.T) {
	store := storage.NewMemoryStorage()
	client := startSMTP(t, config.SMTPConfig{
		Hostname:       "mail.test",
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxMessageSize: 1024,
		MaxConnections: 4,
	}, handlers.Dependencies{Storage: store})

	client.expectCode(t, 220)
	client.command(t, "EHLO test.local")
	client.expectMultiline(t, 250)
	client.command(t, "MAIL FROM:<sender@test.local>")
	client.expectCode(t, 250)
	client.command(t, "RCPT TO:<user@example.com>")
	client.expectCode(t, 250)
	client.command(t, "DATA")
	client.expectCode(t, 354)
	client.write(t, "Subject: test\r\n\r\n.line\r\n.\r\n")
	client.expectCode(t, 250)

	if got := store.MessageCount(); got != 1 {
		t.Fatalf("stored message count = %d, want 1", got)
	}
	if got := store.QueueCount(); got != 1 {
		t.Fatalf("queue item count = %d, want 1", got)
	}
	client.command(t, "QUIT")
	client.expectCode(t, 221)
}

func TestSMTPEnforcesSequenceAndMessageSize(t *testing.T) {
	store := storage.NewMemoryStorage()
	client := startSMTP(t, config.SMTPConfig{
		Hostname:       "mail.test",
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxMessageSize: 16,
		MaxConnections: 4,
	}, handlers.Dependencies{Storage: store})

	client.expectCode(t, 220)
	client.command(t, "RCPT TO:<user@example.com>")
	client.expectCode(t, 503)
	client.command(t, "EHLO test.local")
	client.expectMultiline(t, 250)
	client.command(t, "MAIL FROM:<sender@test.local>")
	client.expectCode(t, 250)
	client.command(t, "RCPT TO:<user@example.com>")
	client.expectCode(t, 250)
	client.command(t, "DATA")
	client.expectCode(t, 354)
	client.write(t, "This message is too large\r\n.\r\n")
	client.expectCode(t, 552)
	if got := store.MessageCount(); got != 0 {
		t.Fatalf("oversized message was stored: count = %d", got)
	}
}

func TestSMTPRejectsHighSpamScore(t *testing.T) {
	client := startSMTP(t, config.SMTPConfig{
		Hostname:       "mail.test",
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxMessageSize: 1024,
		MaxConnections: 4,
		EnableSpam:     true,
	}, handlers.Dependencies{Storage: storage.NewMemoryStorage(), Spam: spamChecker(9)})

	client.expectCode(t, 220)
	client.command(t, "EHLO test.local")
	client.expectMultiline(t, 250)
	client.command(t, "MAIL FROM:<sender@test.local>")
	client.expectCode(t, 250)
	client.command(t, "RCPT TO:<user@example.com>")
	client.expectCode(t, 250)
	client.command(t, "DATA")
	client.expectCode(t, 354)
	client.write(t, "spam content\r\n.\r\n")
	client.expectCode(t, 550)
}

func TestSMTPSessionsAreIsolated(t *testing.T) {
	cfg := config.SMTPConfig{
		Hostname:       "mail.test",
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxMessageSize: 1024,
		MaxConnections: 4,
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		done <- server.NewWithDependencies(cfg, handlers.Dependencies{Storage: storage.NewMemoryStorage()}).Serve(ctx, listener)
	}()
	t.Cleanup(func() {
		cancel()
		_ = listener.Close()
		select {
		case err := <-done:
			if err != nil {
				t.Errorf("SMTP server shutdown: %v", err)
			}
		case <-time.After(2 * time.Second):
			t.Error("SMTP server did not shut down")
		}
	})

	first := dialSMTP(t, listener.Addr().String())
	defer first.conn.Close()
	first.expectCode(t, 220)
	first.command(t, "EHLO first.local")
	first.expectMultiline(t, 250)
	first.command(t, "MAIL FROM:<first@example.test>")
	first.expectCode(t, 250)

	second := dialSMTP(t, listener.Addr().String())
	defer second.conn.Close()
	second.expectCode(t, 220)
	second.command(t, "EHLO second.local")
	second.expectMultiline(t, 250)
	second.command(t, "RCPT TO:<victim@example.test>")
	second.expectCode(t, 503)
}

func TestSMTPRejectsConnectionsAboveLimit(t *testing.T) {
	cfg := config.SMTPConfig{
		Hostname:       "mail.test",
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxMessageSize: 1024,
		MaxConnections: 1,
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		done <- server.NewWithDependencies(cfg, handlers.Dependencies{Storage: storage.NewMemoryStorage()}).Serve(ctx, listener)
	}()
	t.Cleanup(func() {
		cancel()
		_ = listener.Close()
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			t.Error("SMTP server did not shut down")
		}
	})

	first := dialSMTP(t, listener.Addr().String())
	defer first.conn.Close()
	first.expectCode(t, 220)
	second := dialSMTP(t, listener.Addr().String())
	defer second.conn.Close()
	line, err := second.reader.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(line, "421 ") {
		t.Fatalf("connection limit response = %q, want 421", line)
	}
}

func startSMTP(t *testing.T, cfg config.SMTPConfig, deps handlers.Dependencies) *smtpTestClient {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	srv := server.NewWithDependencies(cfg, deps)
	done := make(chan error, 1)
	go func() {
		done <- srv.Serve(ctx, listener)
	}()
	t.Cleanup(func() {
		cancel()
		_ = listener.Close()
		select {
		case err := <-done:
			if err != nil {
				t.Errorf("SMTP server shutdown: %v", err)
			}
		case <-time.After(2 * time.Second):
			t.Error("SMTP server did not shut down")
		}
	})
	return dialSMTP(t, listener.Addr().String())
}

func dialSMTP(t *testing.T, address string) *smtpTestClient {
	t.Helper()
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	_ = conn.SetDeadline(time.Now().Add(2 * time.Second))
	return &smtpTestClient{conn: conn, reader: bufio.NewReader(conn)}
}

type smtpTestClient struct {
	conn   net.Conn
	reader *bufio.Reader
}

func (c *smtpTestClient) command(t *testing.T, line string) {
	t.Helper()
	c.write(t, line+"\r\n")
}

func (c *smtpTestClient) write(t *testing.T, data string) {
	t.Helper()
	if _, err := fmt.Fprint(c.conn, data); err != nil {
		t.Fatal(err)
	}
}

func (c *smtpTestClient) expectCode(t *testing.T, want int) {
	t.Helper()
	line, err := c.reader.ReadString('\n')
	if err != nil {
		t.Fatalf("read SMTP response: %v", err)
	}
	if len(line) < 3 || line[:3] != fmt.Sprintf("%03d", want) {
		t.Fatalf("response = %q, want code %d", line, want)
	}
}

func (c *smtpTestClient) expectMultiline(t *testing.T, want int) {
	t.Helper()
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			t.Fatalf("read multiline SMTP response: %v", err)
		}
		if len(line) < 4 || line[:3] != fmt.Sprintf("%03d", want) {
			t.Fatalf("response = %q, want code %d", line, want)
		}
		if !strings.HasPrefix(line[3:], "-") {
			return
		}
	}
}
