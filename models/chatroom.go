package models

// ChatRoom holds information for the users inside a particular chat room
type ChatRoom struct {
	ID    string   `json:"id"`    // to hold the ID of ChatRoom
	Name  string   `json:"name"`  // to hold the Name of ChatRoom
	Users []string `json:"users"` // to hold array of ID of Users inside the ChatRoom
}
