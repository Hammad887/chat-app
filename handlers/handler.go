package handlers

import (
	runtime "github.com/Hammad887/chat-app"
	"github.com/Hammad887/chat-app/gen/restapi/operations"
)

func NewCustomHandler(api *operations.ChatroomAPI, rt *runtime.Runtime) {
	api.ServiceRegisterUserHandler = RegisterUserHandler(rt)
	api.ServiceLoginHandler = LoginUserHandler(rt)
	api.ServiceLogoutHandler = LogoutUserHandler(rt)
	api.ServiceGetAllChatroomsHandler = GetAllChatroomsHandler(rt)
	api.ServiceGetChatroomHandler = GetChatroomHandler(rt)
	api.ServiceSendMessageHandler = SendMessageHandler(rt)
	api.ServiceGetAllMessagesHandler = GetChatroomMessagesHandler(rt)
	// add all new handlers
}
