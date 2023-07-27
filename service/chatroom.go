package service

import (
	"github.com/Hammad887/chatapp/tree/domain/models"
	"github.com/Hammad887/chatapp/tree/domain/db"
	"github.com/google/uuid"
)

// CreateChatRoom creates a new chat room
func CreateChatRoom(chatRoom *models.ChatRoom) error {
	chatRoom.ID = uuid.New()

	_, err := db.DB.Exec("INSERT INTO chat_rooms (id, name, users) VALUES (?, ?, ?)", chatRoom.ID.String(), chatRoom.Name, chatRoom.Users)

	return err
}

// GetChatRoomByID retrieves a chat room by ID
func GetChatRoomByID(id string) (*models.ChatRoom, error) {
	var chatRoom models.ChatRoom
	err := db.DB.QueryRow("SELECT id, name, users FROM chat_rooms WHERE id = ?", id).Scan(&chatRoom.ID, &chatRoom.Name, &chatRoom.Users)
	return &chatRoom, err
}
