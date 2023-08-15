package models

type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	SenderID  string `json:"sender_id"`
	RoomID    string `json:"room_id"`
	CreatedAt string `json:"created_at"`
}
