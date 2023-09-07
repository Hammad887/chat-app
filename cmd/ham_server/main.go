package main

import (
	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"

	runtime "github.com/Hammad887/chat-app"
	"github.com/Hammad887/chat-app/gen/restapi"
	"github.com/Hammad887/chat-app/gen/restapi/operations"
	handler "github.com/Hammad887/chat-app/handlers"
)

func main() {
	log := logrus.WithField("pkg", "main")

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	rt, err := runtime.NewRuntime()
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server

	api := operations.NewChatroomAPI(swaggerSpec)
	api.Logger = log.Infof

	handler.NewCustomHandler(api, rt)

	server = restapi.NewServer(api)
	server.Port = 8080

	defer server.Shutdown()

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
