// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Hammad887/chat-app/docs/restapi/operations"
	"github.com/Hammad887/chat-app/docs/restapi/operations/service"
)

//go:generate swagger generate server --target ../../docs --name Chatroom --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.ChatroomAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ChatroomAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ServiceLoginHandler == nil {
		api.ServiceLoginHandler = service.LoginHandlerFunc(func(params service.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation service.Login has not yet been implemented")
		})
	}
	if api.ServiceLogoutHandler == nil {
		api.ServiceLogoutHandler = service.LogoutHandlerFunc(func(params service.LogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation service.Logout has not yet been implemented")
		})
	}
	if api.ServiceGetAllChatroomsHandler == nil {
		api.ServiceGetAllChatroomsHandler = service.GetAllChatroomsHandlerFunc(func(params service.GetAllChatroomsParams) middleware.Responder {
			return middleware.NotImplemented("operation service.GetAllChatrooms has not yet been implemented")
		})
	}
	if api.ServiceGetAllMessagesHandler == nil {
		api.ServiceGetAllMessagesHandler = service.GetAllMessagesHandlerFunc(func(params service.GetAllMessagesParams) middleware.Responder {
			return middleware.NotImplemented("operation service.GetAllMessages has not yet been implemented")
		})
	}
	if api.ServiceGetChatroomHandler == nil {
		api.ServiceGetChatroomHandler = service.GetChatroomHandlerFunc(func(params service.GetChatroomParams) middleware.Responder {
			return middleware.NotImplemented("operation service.GetChatroom has not yet been implemented")
		})
	}
	if api.ServiceRegisterUserHandler == nil {
		api.ServiceRegisterUserHandler = service.RegisterUserHandlerFunc(func(params service.RegisterUserParams) middleware.Responder {
			return middleware.NotImplemented("operation service.RegisterUser has not yet been implemented")
		})
	}
	if api.ServiceSendMessageHandler == nil {
		api.ServiceSendMessageHandler = service.SendMessageHandlerFunc(func(params service.SendMessageParams) middleware.Responder {
			return middleware.NotImplemented("operation service.SendMessage has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
