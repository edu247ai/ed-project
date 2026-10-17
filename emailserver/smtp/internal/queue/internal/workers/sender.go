package workers

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"email-server/smtp/internal/dkim"
)

type Sender struct {
	LocalName string
	Timeout   time.Duration
	Signer    *dkim.Signer
}

func NewSender(localName string, timeout time.Duration) *Sender {
	return &Sender{LocalName: localName, Timeout: timeout}
}

func (s *Sender) Send(mxHost, recipient string, raw []byte, mailFrom string) error {
	if s.Signer != nil {
		var err error
		raw, err = s.Signer.SignMessage(raw)
		if err != nil {
			return fmt.Errorf("sign outbound message: %w", err)
		}
	}
	timeout := s.Timeout
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(mxHost, "25"), timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(timeout))
	client := newSMTPClient(conn)
	if err := client.expect(220); err != nil {
		return err
	}
	if err := client.command("EHLO %s", s.LocalName); err != nil {
		return err
	}
	if err := client.expect(250); err != nil {
		return err
	}

	if err := client.command("STARTTLS"); err != nil {
		return err
	}
	code, _, err := client.reply()
	if err != nil {
		return err
	}
	if code == 220 {
		tlsConn := tls.Client(conn, &tls.Config{ServerName: mxHost, MinVersion: tls.VersionTLS12})
		if err := tlsConn.Handshake(); err != nil {
			return fmt.Errorf("SMTP STARTTLS handshake: %w", err)
		}
		conn = tlsConn
		_ = conn.SetDeadline(time.Now().Add(timeout))
		client = newSMTPClient(conn)
		if err := client.command("EHLO %s", s.LocalName); err != nil {
			return err
		}
		if err := client.expect(250); err != nil {
			return err
		}
	} else if code/100 != 5 {
		return fmt.Errorf("unexpected STARTTLS response: %d", code)
	}
	if err := client.command("MAIL FROM:<%s>", mailFrom); err != nil {
		return err
	}
	if err := client.expect(250); err != nil {
		return err
	}
	if err := client.command("RCPT TO:<%s>", recipient); err != nil {
		return err
	}
	code, _, err = client.reply()
	if err != nil {
		return err
	}
	if code != 250 && code != 251 {
		return fmt.Errorf("SMTP recipient rejected with code %d", code)
	}
	if err := client.command("DATA"); err != nil {
		return err
	}
	if err := client.expect(354); err != nil {
		return err
	}
	if err := client.writeMessage(raw); err != nil {
		return err
	}
	if err := client.expect(250); err != nil {
		return err
	}
	_ = client.command("QUIT")
	return nil
}

type smtpClient struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func newSMTPClient(conn net.Conn) *smtpClient {
	return &smtpClient{conn: conn, reader: bufio.NewReader(conn), writer: bufio.NewWriter(conn)}
}

func (c *smtpClient) command(format string, args ...any) error {
	if _, err := fmt.Fprintf(c.writer, format+"\r\n", args...); err != nil {
		return err
	}
	return c.writer.Flush()
}

func (c *smtpClient) reply() (int, string, error) {
	var text strings.Builder
	code := 0
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			return 0, "", err
		}
		if len(line) < 4 {
			return 0, "", fmt.Errorf("malformed SMTP response %q", strings.TrimSpace(line))
		}
		lineCode, err := strconv.Atoi(line[:3])
		if err != nil {
			return 0, "", fmt.Errorf("malformed SMTP response %q", strings.TrimSpace(line))
		}
		if code == 0 {
			code = lineCode
		} else if code != lineCode {
			return 0, "", fmt.Errorf("inconsistent multiline SMTP response")
		}
		text.WriteString(line[4:])
		if line[3] == ' ' {
			return code, strings.TrimSpace(text.String()), nil
		}
		if line[3] != '-' {
			return 0, "", fmt.Errorf("malformed SMTP response separator")
		}
	}
}

func (c *smtpClient) expect(expected int) error {
	code, text, err := c.reply()
	if err != nil {
		return err
	}
	if code != expected {
		return fmt.Errorf("SMTP expected %d, received %d: %s", expected, code, text)
	}
	return nil
}

func (c *smtpClient) writeMessage(raw []byte) error {
	var data bytes.Buffer
	atLineStart := true
	for _, b := range raw {
		if atLineStart && b == '.' {
			data.WriteByte('.')
		}
		data.WriteByte(b)
		atLineStart = b == '\n'
	}
	if data.Len() == 0 || !bytes.HasSuffix(data.Bytes(), []byte("\n")) {
		data.WriteString("\r\n")
	}
	data.WriteString(".\r\n")
	if _, err := io.Copy(c.writer, &data); err != nil {
		return err
	}
	return c.writer.Flush()
}
