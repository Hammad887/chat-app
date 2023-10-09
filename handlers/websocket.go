package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"

	runtime "github.com/Hammad887/chat-app"
	domain "github.com/Hammad887/chat-app/models"
)

const (
	// Request Parameters
	MsgFromParamKey     = "message_from"
	MsgParamKey         = "message"
	RequestTypeParamKey = "requestType"
	RoomIDParamKey      = "room_id"
	SenderIDParamKey    = "sender_id"

	// Request Types
	ConnectToChatRoomKey = "connect_to_chat_room"
	SendMessageKey       = "send_message"
)

// upgrader use for upgrading http connection to websocket connection
var upgrader = websocket.Upgrader{
	// ReadBufferSize specifies the size of the buffer used for reading from the WebSocket connection.
	ReadBufferSize: 1024,

	// WriteBufferSize specifies the size of the buffer used for writing to the WebSocket connection.
	WriteBufferSize: 1024,

	// EnableCompression allows for message compression when set to true.
	EnableCompression: true,
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request, rt *runtime.Runtime, ctx context.Context) error {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer func(Conn *websocket.Conn) error {
		err = Conn.Close()
		if err != nil {
			return err
		}
		return err
	}(conn)

	service := rt.Service()

	for {
		_, requestByte, err := conn.ReadMessage()

		if err != nil {
			// terminate the for loop when connection breaks
			return err
		}

		requestJson := make(map[string]interface{})
		err = json.Unmarshal(requestByte, &requestJson)
		if err != nil {
			return err
		}
		switch requestJson[RequestTypeParamKey] {
		case ConnectToChatRoomKey:
			if _, ok := requestJson[RoomIDParamKey]; ok {
				if !ok {
					msg := "required " + RoomIDParamKey
					if err = WriteResponse(ConnectToChatRoomKey, msg, conn, ctx); err != nil {
						return err
					}
					break
				}
			}

			if _, ok := requestJson[SenderIDParamKey]; ok {
				if !ok {
					msg := "required " + SenderIDParamKey
					if err = WriteResponse(ConnectToChatRoomKey, msg, conn, ctx); err != nil {
						return err
					}
					break
				}
			}

			roomId, ok := requestJson[RoomIDParamKey].(string)
			if !ok {
				msg := "required string type " + RoomIDParamKey
				if err = WriteResponse(ConnectToChatRoomKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			senderId, ok := requestJson[SenderIDParamKey].(string)
			if !ok {
				msg := "required string type " + SenderIDParamKey
				if err = WriteResponse(ConnectToChatRoomKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			row, _ := service.GetChatroom(ctx, roomId)
			if row.ID != roomId {
				msg := "invalid " + RoomIDParamKey
				if err = WriteResponse(ConnectToChatRoomKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			hasUserId := false
			for _, userId := range row.Users {
				if userId == senderId {
					hasUserId = true
				}
			}

			if !hasUserId {
				msg := "invalid " + SenderIDParamKey
				if err = WriteResponse(ConnectToChatRoomKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			if _, ok := domain.WebsocketChatRoomMap[roomId]; ok {
				thisChatRoom := domain.WebsocketChatRoomMap[roomId]
				client := &domain.WebsocketUser{
					ID:         senderId,
					Connection: conn,
				}
				thisChatRoom.ClientWsConnection = append(thisChatRoom.ClientWsConnection, client)
				msg := "Successfully Connect"
				if err = WriteResponse(ConnectToChatRoomKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			thisClient := &domain.WebsocketUser{
				Connection: conn,
				ID:         senderId,
			}

			var clients []*domain.WebsocketUser
			clients = append(clients, thisClient)

			thisChatRoom := &domain.WebsocketChatRoom{
				ID:                 roomId,
				ClientWsConnection: clients,
			}

			domain.WebsocketChatRoomMap[roomId] = thisChatRoom

			msg := "Successfully Connect"
			if err = WriteResponse(ConnectToChatRoomKey, msg, conn, ctx); err != nil {
				return err
			}
			break

		case SendMessageKey:

			if _, ok := requestJson[RoomIDParamKey]; ok {
				if !ok {
					msg := "required " + RoomIDParamKey
					if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
						return err
					}
					break
				}
			}

			if _, ok := requestJson[SenderIDParamKey]; ok {
				if !ok {
					msg := "required " + SenderIDParamKey
					if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
						return err
					}
					break
				}
			}

			if _, ok := requestJson[MsgParamKey]; ok {
				if !ok {
					msg := "required " + MsgParamKey
					if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
						return err
					}
					break
				}
			}

			message, ok := requestJson[MsgParamKey].(string)
			if !ok {
				msg := "required string type " + MsgParamKey
				if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}
			roomId, ok := requestJson[RoomIDParamKey].(string)
			if !ok {
				msg := "required string type " + RoomIDParamKey
				if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}
			senderId, ok := requestJson[SenderIDParamKey].(string)
			if !ok {
				msg := "required string type " + SenderIDParamKey
				if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			row, _ := service.GetChatroom(ctx, roomId)
			if row.ID != roomId {
				msg := "invalid " + RoomIDParamKey
				if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			hasUserId := false
			for _, userId := range row.Users {
				if userId == senderId {
					hasUserId = true
				}
			}

			if !hasUserId {
				msg := "invalid " + SenderIDParamKey
				if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			if _, ok := domain.WebsocketChatRoomMap[roomId]; !ok {
				msg := "invalid " + RoomIDParamKey
				if err = WriteResponse(SendMessageKey, msg, conn, ctx); err != nil {
					return err
				}
				break
			}

			thisChatRoom := domain.WebsocketChatRoomMap[roomId]
			for _, client := range thisChatRoom.ClientWsConnection {
				if client.Connection != conn {
					msg := message
					if err = WriteMessageWithSenderId(SendMessageKey, msg, senderId, client.Connection, ctx); err != nil {
						return err
					}
					break
				}
			}

		default:
			log(ctx).Debugln("Unknown Method: ", requestJson[RequestTypeParamKey])
			msg := "Unknown Method: " + RequestTypeParamKey
			if err = WriteResponse(RequestTypeParamKey, msg, conn, ctx); err != nil {
				return err
			}
			break
		}
	}
}

// WriteResponse sends a message back to the client through an established WebSocket connection.
func WriteResponse(requestTypeValue string, message string, conn *websocket.Conn, ctx context.Context) error {
	Response := make(map[string]interface{})
	Response[RequestTypeParamKey] = requestTypeValue
	Response[MsgParamKey] = message
	responseBytes, _ := json.Marshal(Response)
	err := conn.WriteMessage(websocket.TextMessage, responseBytes)
	if err != nil {
		log(ctx).Error("error in write on websocket: ", err)
		return err
	}
	return err
}

// WriteMessageWithSenderId sends a message back to the client through an established WebSocket connection with sender id.
func WriteMessageWithSenderId(requestTypeValue string, message string, senderId string, conn *websocket.Conn, ctx context.Context) error {
	Response := make(map[string]interface{})
	Response[RequestTypeParamKey] = requestTypeValue
	Response[MsgParamKey] = message
	Response[MsgFromParamKey] = senderId
	responseBytes, _ := json.Marshal(Response)
	err := conn.WriteMessage(websocket.TextMessage, responseBytes)
	if err != nil {
		log(ctx).Error("error in write on websocket: ", err)
		return err
	}
	return err
}
