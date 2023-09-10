package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	runtime "github.com/Hammad887/chat-app"
	domainErr "github.com/Hammad887/chat-app/errors"
	docsModel "github.com/Hammad887/chat-app/gen/models"
	"github.com/Hammad887/chat-app/gen/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type getChatroomMessages struct {
	rt *runtime.Runtime
}

// Handle implements service.RegisterUserHandler.
func (r *getChatroomMessages) Handle(params service.GetAllMessagesParams) middleware.Responder {
	ctx := context.Background()
	uuid := params.ChatroomID
	messages, err := r.rt.Service().GetChatroomMessages(context.Background(), uuid)
	if errors.Is(err, domainErr.ErrConflict) {
		log(context.Background()).Errorf("user with given email is already exist in database", err)

		return service.NewGetAllMessagesNotFound().WithPayload(&docsModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusConflict)),
			Message: swag.String(err.Error()),
		})
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewGetAllMessagesDefault(http.StatusInternalServerError).WithPayload(&docsModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("got messages %v", messages)
	return service.NewGetAllMessagesOK().WithPayload(asMessagesResponse(messages))

}

// GetChatroomMessagesHandler returns a handler for retrieving messages from a chat room.
func GetChatroomMessagesHandler(rt *runtime.Runtime) service.GetAllMessagesHandler {
	return &getChatroomMessages{rt: rt}
}
