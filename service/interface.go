package service

import (
	"context"

	"github.com/Hammad887/chat-app/db"
	"github.com/Hammad887/chat-app/models"
)

// Manager defines the available functions for the given service implementation.
type Manager interface {
	GetChatroom(ctx context.Context, id string) (*models.ChatRoom, error)
	GetChatroomMessages(ctx context.Context, id string) ([]*models.Message, error)
	GetChatrooms(ctx context.Context) ([]*models.ChatRoom, error)
	RegisterUser(ctx context.Context, user *models.User) (bool, error)
	LoginUser(ctx context.Context, email string, password string) (string, error)
	LogoutUser(ctx context.Context, token string) (bool, error)
	SendMessage(ctx context.Context, id string, message *models.Message) error
}

type service struct {
	db db.DataStore
}

// NewService return new service object
func NewService(store *db.DataStore) Manager {
	return &service{db: *store}
}
