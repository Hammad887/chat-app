package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	wraperrors "github.com/pkg/errors"

	genModel "github.com/Hammad887/chat-app/docs/models"
	"github.com/Hammad887/chat-app/models"
)

// UpdateUser update user's name and address
func (s *service) GetChatroom(ctx context.Context, id string) (*genModel.Chatroom, error) {
	chatroom, err := s.db.GetChatroom(ctx, id)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return asChatroomResponse(chatroom), nil
}

// GetChatroomMessages implements Manager.
func (s *service) GetChatroomMessages(ctx context.Context, id string) ([]*genModel.Message, error) {
	messages, err := s.db.GetChatroomMessages(ctx, id)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return asMessagesResponse(messages), nil
}

// GetChatrooms implements Manager.
func (s *service) GetChatrooms(ctx context.Context) ([]*genModel.Chatroom, error) {
	chatrooms, err := s.db.GetChatrooms(ctx)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return asChatroomsResponse(*chatrooms), nil
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
func (s *service) RegisterUser(ctx context.Context, user *genModel.User) (bool, error) {
	success, err := s.db.RegisterUser(ctx, &models.User{
		Name:     *user.Name,
		Email:    user.Email.String(),
		Password: *user.Password,
	})
	if err != nil {
		return false, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return success, nil
}

// SendMessage implements Manager.
func (s *service) SendMessage(ctx context.Context, id string, message *genModel.Message) (bool, error) {
	success, err := s.db.SendMessage(ctx, id, &models.Message{
		CreatedAt: time.Now().String(),
		ID:        uuid.NewString(),
		RoomID:    *message.RoomID,
		SenderID:  *message.SenderID,
		Text:      *message.Text,
	})
	if err != nil {
		return false, wraperrors.Wrap(err, "failed to find user with given id")
	}

	return success, nil
}
