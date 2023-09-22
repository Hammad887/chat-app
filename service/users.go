package service

import (
	"context"
	"fmt"

	"github.com/Hammad887/chat-app/models"
)

// UpdateUser update user's name and address
func (s *service) GetChatroom(ctx context.Context, id string) (*models.ChatRoom, error) {
	chatroom, err := s.db.GetChatroom(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with given id: %w", err)
	}

	return chatroom, nil
}

// GetChatroomMessages implements Manager.
func (s *service) GetChatroomMessages(ctx context.Context, id string) ([]*models.Message, error) {
	messages, err := s.db.GetChatroomMessages(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with given id: %w", err)
	}

	return messages, nil
}

// ListChatRoom implements Manager.
func (s *service) ListChatRoom(ctx context.Context) ([]*models.ChatRoom, error) {
	chatrooms, err := s.db.ListChatRoom(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with given id: %w", err)
	}

	return chatrooms, nil
}

// LoginUser implements Manager.
func (s *service) LoginUser(ctx context.Context, email string, password string) (string, error) {
	token, err := s.db.LoginUser(ctx, email, password)
	if err != nil {
		return "", fmt.Errorf("failed to find user with given id: %w", err)
	}

	return token, nil
}

// LogoutUser implements Manager.
func (s *service) LogoutUser(ctx context.Context, token string) (bool, error) {
	success, err := s.db.LogoutUser(ctx, token)
	if err != nil {
		return false, fmt.Errorf("failed to find user with given id: %w", err)
	}

	return success, nil
}

// RegisterUser implements Manager.
func (s *service) RegisterUser(ctx context.Context, user *models.User) (bool, error) {
	success, err := s.db.RegisterUser(ctx, user)
	if err != nil {
		return false, fmt.Errorf("failed to find user with given id: %w", err)
	}

	return success, nil
}

// SendMessage implements Manager.
func (s *service) SendMessage(ctx context.Context, id string, message *models.Message) error {
	err := s.db.SendMessage(ctx, id, message)
	if err != nil {
		return fmt.Errorf("failed to find user with given id: %w", err)
	}

	return nil
}
