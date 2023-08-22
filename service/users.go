package service

import (
	"context"

	wraperrors "github.com/pkg/errors"

	"github.com/Hammad887/chat-app/models"
)

// UpdateUser update user's name and address
func (s *service) GetChatroom(ctx context.Context, id string) (*models.ChatRoom, error) {
	chatroom, err := s.db.GetChatroom(ctx, id)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return chatroom, nil
}

// GetChatroomMessages implements Manager.
func (s *service) GetChatroomMessages(ctx context.Context, id string) ([]*models.Message, error) {
	messages, err := s.db.GetChatroomMessages(ctx, id)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return messages, nil
}

// GetChatrooms implements Manager.
func (s *service) GetChatrooms(ctx context.Context) ([]*models.ChatRoom, error) {
	chatrooms, err := s.db.GetChatrooms(ctx)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return chatrooms, nil
}

// LoginUser implements Manager.
func (s *service) LoginUser(ctx context.Context, email string, password string) (string, error) {
	token, err := s.db.LoginUser(ctx, email, password)
	if err != nil {
		return "", wraperrors.Wrap(err, "failed to find user with given id")
	}

	return token, nil
}

// LogoutUser implements Manager.
func (s *service) LogoutUser(ctx context.Context, token string) (bool, error) {
	success, err := s.db.LogoutUser(ctx, token)
	if err != nil {
		return false, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return success, nil
}

// RegisterUser implements Manager.
func (s *service) RegisterUser(ctx context.Context, user *models.User) (bool, error) {
	success, err := s.db.RegisterUser(ctx, user)
	if err != nil {
		return false, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return success, nil
}

// SendMessage implements Manager.
func (s *service) SendMessage(ctx context.Context, id string, message *models.Message) error {
	err := s.db.SendMessage(ctx, id, message)
	if err != nil {
		return wraperrors.Wrap(err, "failed to find user with given id")
	}

	return nil
}
