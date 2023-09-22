package handlers

import (
	"context"
	"errors"

	runtime "github.com/Hammad887/chat-app"
	domainErr "github.com/Hammad887/chat-app/errors"
	docsModel "github.com/Hammad887/chat-app/gen/models"
	"github.com/Hammad887/chat-app/gen/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
)

type logoutUser struct {
	rt *runtime.Runtime
}

// Handle implements service.RegisterUserHandler.
func (r *logoutUser) Handle(params service.LogoutParams) middleware.Responder {
	ctx := context.Background()

	token := params.Logout

	success, err := r.rt.Service().LogoutUser(context.Background(), *token.Token)
	if errors.Is(err, domainErr.ErrConflict) {
		log(context.Background()).Errorf("user with given email is already exist in database", err)

		return service.NewLogoutBadRequest()
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewLogoutInternalServerError()
	}

	log(ctx).Infof("got token %v", token)
	return service.NewLogoutOK().WithPayload(&docsModel.LogoutSuccess{
		Success: success,
		Token:   *token.Token,
	})
}

// LogoutUserHandler creates and returns a handler for logging out users using the provided runtime.
func LogoutUserHandler(rt *runtime.Runtime) service.LogoutHandler {
	return &logoutUser{rt: rt}
}
