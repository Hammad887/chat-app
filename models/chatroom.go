package models

type ChatRoom struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}
