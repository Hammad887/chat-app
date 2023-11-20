package models

import "github.com/gorilla/websocket"

// WebsocketUser holds information for a WebsocketUser user
type WebsocketUser struct {
	ID         string          `json:"id"`
	Connection *websocket.Conn `json:"connection,omitempty"`
}
