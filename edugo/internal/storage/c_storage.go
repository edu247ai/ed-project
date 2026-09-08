package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"time"
)

type CStorage struct {
	BinaryPath string
	Timeout    time.Duration
}

func NewCStorage() *CStorage {
	return &CStorage{
		BinaryPath: "./c_storage_binary",
		Timeout:    10 * time.Second, // prevent deadlocks
	}
}

func (c *CStorage) Save(path string, data io.Reader) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, c.BinaryPath, "save", path)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to open stdin: %w", err)
	}

	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start C binary: %w", err)
	}

	go func() {
		io.Copy(stdin, data)
		stdin.Close()
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("C storage save failed: %v | stderr: %s", err, stderr.String())
	}

	return nil
}

func (c *CStorage) Get(path string) (io.ReadCloser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, c.BinaryPath, "get", path)

	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to open stdout: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start C binary: %w", err)
	}

	return stdout, nil
}

func (c *CStorage) Delete(path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, c.BinaryPath, "delete", path)

	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("C storage delete failed: %v | stderr: %s", err, stderr.String())
	}

	return nil
}

func (c *CStorage) PresignURL(path string, expiry time.Duration) (string, error) {
	return "", fmt.Errorf("C storage does not support presigned URLs")
}
