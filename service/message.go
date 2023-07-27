package service

import (
	"github.com/Hammad887/chatapp/models"
	"github.com/Hammad887/chatapp/db"
	"github.com/google/uuid"
)

// CreateMessage creates a new message in a chat room
func CreateMessage(message *models.Message) error {
	message.ID = uuid.New()

	_, err := db.DB.Exec("INSERT INTO messages (id, text, sender_id, room_id, created_at) VALUES (?, ?, ?, ?, ?)", message.ID.String(), message.Text, message.SenderID, message.RoomID, message.CreatedAt)

	return err
}

// GetMessagesByRoomID retrieves messages by chat room ID
func GetMessagesByRoomID(roomID string) ([]*models.Message, error) {
	rows, err := db.DB.Query("SELECT id, text, sender_id, room_id, created_at FROM messages WHERE room_id = ?", roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.Text, &message.SenderID, &message.RoomID, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}
