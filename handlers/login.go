package handlers

import (
	"context"
	"errors"

	runtime "github.com/Hammad887/chat-app"
	domainErr "github.com/Hammad887/chat-app/errors"
	genModel "github.com/Hammad887/chat-app/gen/models"
	"github.com/Hammad887/chat-app/gen/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
)

type loginUser struct {
	rt *runtime.Runtime
}

// Handle implements service.RegisterUserHandler.
func (r *loginUser) Handle(params service.LoginParams) middleware.Responder {
	ctx := context.Background()

	email := params.Login.Email
	password := params.Login.Password

	token, err := r.rt.Service().LoginUser(context.Background(), *email, *password)
	if errors.Is(err, domainErr.ErrConflict) {
		log(context.Background()).Errorf("user with given email is already exist in database", err)

		return service.NewLoginBadRequest()
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewLoginInternalServerError()
	}

	log(ctx).Infof("got token %v", token)
	return service.NewLoginOK().WithPayload(&genModel.LoginSuccess{
		Token:   token,
		Success: true,
	})
}

// LoginUserHandler returns a handler that manages user login.
func LoginUserHandler(rt *runtime.Runtime) service.LoginHandler {
	return &loginUser{rt: rt}
}
