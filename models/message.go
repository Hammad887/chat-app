package models

import (
	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	SenderID  uuid.UUID `json:"sender_id"`
	RoomID    uuid.UUID `json:"room_id"`
	CreatedAt string    `json:"created_at"`
}
