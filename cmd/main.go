package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Hammad887/chatapp/db"
	"github.com/Hammad887/chatapp/handlers"
)

func main() {
	// Initialize the database connection
	db.InitializeDB()
	defer db.DB.Close()

	router := mux.NewRouter().StrictSlash(true)

	public := router.PathPrefix("/api").Subrouter()
	public.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	port := ":8000"

	fmt.Printf("Server started on port %s\n", port)
	http.ListenAndServe(port, router)
}
