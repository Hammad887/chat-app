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

type sendMessage struct {
	rt *runtime.Runtime
}

// Handle implements service.RegisterUserHandler.
func (r *sendMessage) Handle(params service.SendMessageParams) middleware.Responder {
	ctx := context.Background()

	message := params.Message
	id := params.ChatroomID

	success, err := r.rt.Service().SendMessage(context.Background(), id, message)
	if errors.Is(err, domainErr.ErrConflict) {
		log(context.Background()).Errorf("user with given email is already exist in database", err)

		return service.NewSendMessageUnauthorized().WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusConflict)),
			Message: swag.String(err.Error()),
		})
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewSendMessageDefault(http.StatusInternalServerError).WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("sent messages %v", success)
	return service.NewSendMessageCreated().WithPayload(&genModel.SuccessResponse{
		Success: success,
	})
}

func SendMessageHandler(rt *runtime.Runtime) service.SendMessageHandler {
	return &sendMessage{rt: rt}
}
