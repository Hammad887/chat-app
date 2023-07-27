package models

import (
	"github.com/google/uuid"
)

type ChatRoom struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Users []string  `json:"users"`
}
