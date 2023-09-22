package handlers

import (
	genModel "github.com/Hammad887/chat-app/gen/models"
	domain "github.com/Hammad887/chat-app/models"
)

func asChatroomResponse(chatroom *domain.ChatRoom) *genModel.Chatroom {
	return &genModel.Chatroom{
		ID:    &chatroom.ID,
		Name:  &chatroom.Name,
		Users: chatroom.Users,
	}
}

func asChatroomsResponse(chatrooms []*domain.ChatRoom) []*genModel.Chatroom {
	returnedChatrooms := make([]*genModel.Chatroom, 0)

	for _, chatroom := range chatrooms {
		returnedChatrooms = append(returnedChatrooms, &genModel.Chatroom{
			ID:    &chatroom.ID,
			Name:  &chatroom.Name,
			Users: chatroom.Users,
		})
	}
	return returnedChatrooms
}

func asMessagesResponse(messages []*domain.Message) []*genModel.Message {
	returnedMessages := make([]*genModel.Message, 0)

	for _, message := range messages {
		returnedMessages = append(returnedMessages, &genModel.Message{
			CreatedAt: message.CreatedAt,
			ID:        message.ID,
			SenderID:  &message.SenderID,
			RoomID:    &message.RoomID,
			Text:      &message.Text,
		})
	}
	return returnedMessages
}
