package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChatRoom struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

type Message struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	SenderID  int       `json:"sender_id"`
	RoomID    int       `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB
var users []User
var chatRooms []ChatRoom
var messages []Message
var jwtSecretKey = []byte("hello")
var revokedTokens []string

func main() {
	users = make([]User, 0)
	chatRooms = make([]ChatRoom, 0)
	messages = make([]Message, 0)
	revokedTokens = make([]string, 0)

	// Initialize the database connection
	var err error
	db, err = sql.Open("mysql", "hammad:Hammad_887@/chatapp")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	public := router.PathPrefix("/api").Subrouter()
	public.HandleFunc("/register", registerHandler).Methods("POST")
	public.HandleFunc("/login", loginHandler).Methods("POST")

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(tokenVerificationMiddleware)
	protected.HandleFunc("/logout", logoutHandler).Methods("POST")
	protected.HandleFunc("/chat/rooms", chatRoomsHandler).Methods("GET")
	protected.HandleFunc("/chat/rooms/{id}", chatRoomHandler).Methods("GET")
	protected.HandleFunc("/chat/rooms/{id}/messages", sendMessageHandler).Methods("POST")
	protected.HandleFunc("/chat/rooms/{id}/messages", chatRoomMessagesHandler).Methods("GET")
	protected.HandleFunc("/chat/rooms/{roomID}/assign/{userID}", assignUserToChatRoom).Methods("POST")

	port := ":8000"

	fmt.Printf("Server started on port %s\n", port)
	http.ListenAndServe(port, router)
}

func extractTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) == 2 {
			return splitToken[1]
		}
	}
	return ""
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var storedUser User
	if err := db.QueryRow("SELECT * FROM users WHERE email = ?", user.Email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check password

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    storedUser.ID,
		"name":  storedUser.Name,
		"email": storedUser.Email,
	})

	tokenString, _ := token.SignedString(jwtSecretKey)

	w.Write([]byte(tokenString))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := extractTokenFromRequest(r)

	revokedTokens = append(revokedTokens, tokenString)

	w.WriteHeader(http.StatusOK)
}

func tokenVerificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromRequest(r)

		// Check if token is in revoked tokens
		for _, revokedToken := range revokedTokens {
			if revokedToken == tokenString {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func chatRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM chatrooms")
	if err != nil {
		http.Error(w, "Failed to retrieve chatrooms", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	chatRooms := make([]ChatRoom, 0)
	for rows.Next() {
		var chatRoom ChatRoom
		if err := rows.Scan(&chatRoom.ID, &chatRoom.Name); err != nil {
			http.Error(w, "Failed to retrieve chatroom", http.StatusInternalServerError)
			return
		}

		chatRooms = append(chatRooms, chatRoom)
	}

	json.NewEncoder(w).Encode(chatRooms)
}

func chatRoomHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var chatRoom ChatRoom
	if err := db.QueryRow("SELECT * FROM chatrooms WHERE id = ?", id).Scan(&chatRoom.ID, &chatRoom.Name); err != nil {
		http.Error(w, "Chatroom not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(chatRoom)
}

func assignUserToChatRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID, ok := vars["roomID"]
	if !ok {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	userID, ok := vars["userID"]
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO room_user (room_id, user_id) VALUES (?, ?)", roomID, userID)
	if err != nil {
		http.Error(w, "Failed to assign user to chatroom", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO messages (text, sender_id, room_id, created_at) VALUES (?, ?, ?, ?)", message.Text, message.SenderID, message.RoomID, time.Now())
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func chatRoomMessagesHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	rows, err := db.Query("SELECT * FROM messages WHERE room_id = ?", id)
	if err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	messages := make([]Message, 0)
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.Text, &message.SenderID, &message.RoomID, &message.CreatedAt); err != nil {
			http.Error(w, "Failed to retrieve message", http.StatusInternalServerError)
			return
		}

		messages = append(messages, message)
	}

	json.NewEncoder(w).Encode(messages)
}
