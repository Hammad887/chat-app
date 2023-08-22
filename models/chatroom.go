package models

// ChatRoom holds information for the users inside a particular chat room
type ChatRoom struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}
