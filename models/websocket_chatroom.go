package models

// WebsocketChatRoom holds information for the users inside a particular chat room
type WebsocketChatRoom struct {
	ID                 string           `json:"id"`                             // to hold the ID of ChatRoom
	ClientWsConnection []*WebsocketUser `json:"client_ws_connection,omitempty"` // to hold array of User's complete struct information with websocket connection address
}

// WebsocketChatRoomMap holds *WebsocketChatRoom object as value, and It's ID as key
var WebsocketChatRoomMap = make(map[string]*WebsocketChatRoom)
