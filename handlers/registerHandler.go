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

type registerUser struct {
	rt *runtime.Runtime
}

// Handle implements service.RegisterUserHandler.
func (r *registerUser) Handle(params service.RegisterUserParams) middleware.Responder {
	ctx := context.Background()

	user := params.User

	success, err := r.rt.Service().RegisterUser(context.Background(), &domain.User{
		ID:       *user.ID,
		Name:     *user.Name,
		Email:    user.Email.String(),
		Password: *user.Password,
	})
	if errors.Is(err, domainErr.ErrConflict) {
		log(context.Background()).Errorf("user with given email is already exist in database", err)

		return service.NewRegisterUserConflict().WithPayload(&docsModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusConflict)),
			Message: swag.String(err.Error()),
		})
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewRegisterUserDefault(http.StatusInternalServerError).WithPayload(&docsModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("created user %v", success)
	return service.NewRegisterUserCreated().WithPayload(&docsModel.SuccessResponse{
		Success: success,
		Message: "User registered successfully.",
	})
}

// RegisterUserHandler creates and returns a handler for registering users using the provided runtime.
func RegisterUserHandler(rt *runtime.Runtime) service.RegisterUserHandler {
	return &registerUser{rt: rt}
}
