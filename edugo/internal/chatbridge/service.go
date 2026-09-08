package chatbridge

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
	ws   *WebSocketBridge
}

func NewService(repo *Repository, ws *WebSocketBridge) *Service {
	return &Service{repo: repo, ws: ws}
}

func (s *Service) SendMessage(ctx context.Context, roomID, userID int, content string) (*ChatMessage, error) {
	msg := &ChatMessage{
		RoomID:  roomID,
		UserID:  userID,
		Content: content,
	}

	// Save to DB
	if err := s.repo.SaveMessage(ctx, msg); err != nil {
		return nil, err
	}

	// Forward to C WebSocket server
	if err := s.ws.Broadcast(roomID, userID, content); err != nil {
		return nil, fmt.Errorf("failed to broadcast: %w", err)
	}

	return msg, nil
}

func (s *Service) GetRoomMessages(ctx context.Context, roomID int) ([]ChatMessage, error) {
	return s.repo.ListMessages(ctx, roomID)
}
