package chatbridge

import (
	"bytes"
	"fmt"
	"os/exec"
)

type WebSocketBridge struct {
	Binary string
}

func NewWebSocketBridge() *WebSocketBridge {
	return &WebSocketBridge{
		Binary: "./chat-c/server", // your C WebSocket server
	}
}

func (w *WebSocketBridge) Broadcast(roomID, userID int, content string) error {
	cmd := exec.Command(w.Binary, "broadcast", fmt.Sprint(roomID), fmt.Sprint(userID))

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start error: %w", err)
	}

	// Send message content to C server
	stdin.Write([]byte(content))
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("broadcast failed: %v | %s", err, stderr.String())
	}

	return nil
}
