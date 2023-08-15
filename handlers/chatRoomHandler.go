package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	runtime "github.com/Hammad887/chat-app"
	genModel "github.com/Hammad887/chat-app/docs/models"
	"github.com/Hammad887/chat-app/docs/restapi/operations/service"
	domainErr "github.com/Hammad887/chat-app/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type getChatroom struct {
	rt *runtime.Runtime
}

// Handle implements service.RegisterUserHandler.
func (r *getChatroom) Handle(params service.GetChatroomParams) middleware.Responder {
	ctx := context.Background()
	uuid := params.ChatroomID
	chatroom, err := r.rt.Service().GetChatroom(context.Background(), uuid)
	if errors.Is(err, domainErr.ErrConflict) {
		log(context.Background()).Errorf("user with given email is already exist in database", err)

		return service.NewGetChatroomNotFound().WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusConflict)),
			Message: swag.String(err.Error()),
		})
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewGetChatroomDefault(http.StatusInternalServerError).WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("got chatroom id %v", chatroom)
	return service.NewGetChatroomOK().WithPayload(chatroom)

}

func GetChatroomHandler(rt *runtime.Runtime) service.GetChatroomHandler {
	return &getChatroom{rt: rt}
}
