package service

import (
	"context"

	"github.com/Hammad887/chat-app/db"
	genModel "github.com/Hammad887/chat-app/docs/models"
)

// Manager defines the available functions for the given service implementation.
type Manager interface {
	GetChatroom(ctx context.Context, id string) (*genModel.Chatroom, error)
	GetChatroomMessages(ctx context.Context, id string) ([]*genModel.Message, error)
	GetChatrooms(ctx context.Context) ([]*genModel.Chatroom, error)
	RegisterUser(ctx context.Context, user *genModel.User) (bool, error)
	LoginUser(ctx context.Context, email string, password string) (string, error)
	LogoutUser(ctx context.Context, token string) (bool, error)
	SendMessage(ctx context.Context, id string, message *genModel.Message) (bool, error)
}

type service struct {
	db db.DataStore
}

// NewService return new service object
func NewService(store *db.DataStore) Manager {
	return &service{db: *store}
}
