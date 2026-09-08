package chatbridge

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Client struct {
	conn   net.Conn
	reader *bufio.Reader
	addr   string
}

func New(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to chat server: %w", err)
	}

	return &Client{
		conn:   conn,
		reader: bufio.NewReader(conn),
		addr:   addr,
	}, nil
}

// Reconnect if connection drops
func (c *Client) reconnect() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	c.reader = bufio.NewReader(conn)
	return nil
}

// Send a message with newline framing
func (c *Client) Send(msg string) error {
	_, err := c.conn.Write([]byte(msg + "\n"))
	if err != nil {
		// Try reconnect once
		if recErr := c.reconnect(); recErr != nil {
			return fmt.Errorf("send failed and reconnect failed: %v | %v", err, recErr)
		}
		// Retry send
		_, err = c.conn.Write([]byte(msg + "\n"))
	}
	return err
}

// Receive one line (blocking)
func (c *Client) Receive() (string, error) {
	line, err := c.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return line, nil
}

// Non-blocking receive with timeout
func (c *Client) ReceiveWithTimeout(timeout time.Duration) (string, error) {
	c.conn.SetReadDeadline(time.Now().Add(timeout))
	line, err := c.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return line, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
