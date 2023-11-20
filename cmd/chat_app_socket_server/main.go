// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	request "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	runtime "github.com/Hammad887/chat-app"
	"github.com/Hammad887/chat-app/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	router := mux.NewRouter()
	// Set up CORS handling
	headers := request.AllowedHeaders([]string{"*"})
	origins := request.AllowedOrigins([]string{"*"}) // Allow requests from any origin

	// Add CORS middleware
	router.Use(request.CORS(headers, origins))

	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	//service := rt.Service()

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err = handlers.WebSocketHandler(w, r, rt, ctx)
		if err != nil {
			log.Println(err)
		}
	})

	port := os.Getenv("SERVER_PORT")

	fmt.Printf("Server is running on port %s\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
