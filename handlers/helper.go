package handlers

import (
	docsModel "github.com/Hammad887/chat-app/gen/models"
	domain "github.com/Hammad887/chat-app/models"
)

func asChatroomResponse(chatroom *domain.ChatRoom) *docsModel.Chatroom {
	return &docsModel.Chatroom{
		ID:    &chatroom.ID,
		Name:  &chatroom.Name,
		Users: chatroom.Users,
	}
}

func asChatroomsResponse(chatrooms []*domain.ChatRoom) []*docsModel.Chatroom {
	var returnedChatrooms []*docsModel.Chatroom

	for _, chatroom := range chatrooms {
		returnedChatrooms = append(returnedChatrooms, &docsModel.Chatroom{
			ID:    &chatroom.ID,
			Name:  &chatroom.Name,
			Users: chatroom.Users,
		})
	}
	return returnedChatrooms
}

func asMessagesResponse(messages []*domain.Message) []*docsModel.Message {
	var returnedMessages []*docsModel.Message

	for _, message := range messages {
		returnedMessages = append(returnedMessages, &docsModel.Message{
			CreatedAt: message.CreatedAt,
			ID:        message.ID,
			SenderID:  &message.SenderID,
			RoomID:    &message.RoomID,
			Text:      &message.Text,
		})
	}
	return returnedMessages
}
