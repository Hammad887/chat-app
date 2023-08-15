package db

import (
	"context"
	"log"

	"github.com/Hammad887/chat-app/models"
)

// DataStore is an interface for query ops
type DataStore interface {
	GetChatroom(ctx context.Context, id string) (*models.ChatRoom, error)
	GetChatroomMessages(ctx context.Context, id string) ([]*models.Message, error)
	GetChatrooms(ctx context.Context) (*[]models.ChatRoom, error)
	RegisterUser(ctx context.Context, user *models.User) (bool, error)
	LoginUser(ctx context.Context, email string, password string) (string, error)
	LogoutUser(ctx context.Context, token string) (bool, error)
	SendMessage(ctx context.Context, id string, message *models.Message) (bool, error)
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	if _, ok := datastoreFactories[name]; ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
