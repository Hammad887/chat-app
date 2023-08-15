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

type getChatrooms struct {
	rt *runtime.Runtime
}

// Handle implements service.RegisterUserHandler.
func (r *getChatrooms) Handle(params service.GetAllChatroomsParams) middleware.Responder {
	ctx := context.Background()
	chatrooms, err := r.rt.Service().GetChatrooms(ctx)
	if errors.Is(err, domainErr.ErrConflict) {
		log(context.Background()).Errorf("user with given email is already exist in database", err)

		return service.NewGetAllChatroomsNotFound().WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusConflict)),
			Message: swag.String(err.Error()),
		})
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewGetAllChatroomsDefault(http.StatusInternalServerError).WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("got chatrooms %v", chatrooms)
	return service.NewGetAllChatroomsOK().WithPayload(chatrooms)

}

func GetAllChatroomsHandler(rt *runtime.Runtime) service.GetAllChatroomsHandler {
	return &getChatrooms{rt: rt}
}
