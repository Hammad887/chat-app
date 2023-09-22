package db

import (
	"context"

	"github.com/Hammad887/chat-app/models"
)

// DataStore is an interface for query ops
type DataStore interface {
	GetChatroom(ctx context.Context, id string) (*models.ChatRoom, error)
	GetChatroomMessages(ctx context.Context, id string) ([]*models.Message, error)
	ListChatRoom(ctx context.Context) ([]*models.ChatRoom, error)
	RegisterUser(ctx context.Context, user *models.User) (bool, error)
	LoginUser(ctx context.Context, email string, password string) (string, error)
	LogoutUser(ctx context.Context, token string) (bool, error)
	SaveMessage(ctx context.Context, id string, message *models.Message) error
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}
