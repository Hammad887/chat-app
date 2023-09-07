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
	domain "github.com/Hammad887/chat-app/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type sendMessage struct {
	rt *runtime.Runtime
}

// Handle implements service.RegisterUserHandler.
func (r *sendMessage) Handle(params service.SendMessageParams) middleware.Responder {
	ctx := context.Background()

	message := params.Message
	id := params.ChatroomID

	err := r.rt.Service().SendMessage(context.Background(), id, &domain.Message{
		ID:        message.ID,
		Text:      *message.Text,
		SenderID:  *message.SenderID,
		RoomID:    *message.RoomID,
		CreatedAt: message.CreatedAt,
	})
	if errors.Is(err, domainErr.ErrConflict) {
		log(context.Background()).Errorf("user with given email is already exist in database", err)

		return service.NewSendMessageUnauthorized().WithPayload(&docsModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusConflict)),
			Message: swag.String(err.Error()),
		})
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewSendMessageDefault(http.StatusInternalServerError).WithPayload(&docsModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("sent messages %v", true)
	return service.NewSendMessageCreated().WithPayload(&docsModel.SuccessResponse{
		Success: true,
	})
}

func SendMessageHandler(rt *runtime.Runtime) service.SendMessageHandler {
	return &sendMessage{rt: rt}
}
