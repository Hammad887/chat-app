package service

import (
	"context"
	"fmt"

	"github.com/Hammad887/chat-app/models"
)

// GetChatroom retrieves a chatroom by its ID.
func (s *service) GetChatroom(ctx context.Context, id string) (*models.ChatRoom, error) {
	chatroom, err := s.db.GetChatroom(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with given id: %w", err)
	}

	return chatroom, nil
}

// GetChatroomMessages retrieves messages for a given chatroom ID.
func (s *service) GetChatroomMessages(ctx context.Context, id string) ([]*models.Message, error) {
	messages, err := s.db.GetChatroomMessages(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve messages for the chatroom with given id: %w", err)
	}

	return messages, nil
}

// ListChatRoom lists all available chat rooms.
func (s *service) ListChatRoom(ctx context.Context) ([]*models.ChatRoom, error) {
	chatrooms, err := s.db.ListChatRoom(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list chat rooms from the database: %w", err)
	}

	return chatrooms, nil
}

// LoginUser handles user login and returns a JWT token.
func (s *service) LoginUser(ctx context.Context, email string, password string) (string, error) {
	token, err := s.db.LoginUser(ctx, email, password)
	if err != nil {
		return "", fmt.Errorf("failed to login user with given email: %w", err)
	}

	return token, nil
}

// LogoutUser logs out a user, revoking their token.
func (s *service) LogoutUser(ctx context.Context, token string) (bool, error) {
	success, err := s.db.LogoutUser(ctx, token)
	if err != nil {
		return false, fmt.Errorf("failed to logout user with given token: %w", err)
	}

	return success, nil
}

// RegisterUser registers a new user.
func (s *service) RegisterUser(ctx context.Context, user *models.User) (bool, error) {
	success, err := s.db.RegisterUser(ctx, user)
	if err != nil {
		return false, fmt.Errorf("failed to register user: %w", err)
	}

	return success, nil
}

// SendMessage sends a message to a chatroom.
func (s *service) SendMessage(ctx context.Context, id string, message *models.Message) error {
	err := s.db.SendMessage(ctx, id, message)
	if err != nil {
		return fmt.Errorf("failed to send message to the chatroom with given id: %w", err)
	}

	return nil
}
