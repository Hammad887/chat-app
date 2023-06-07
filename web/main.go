package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

type ChatRoom struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Users []string  `json:"users"`
}

type Message struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	SenderID  uuid.UUID `json:"sender_id"`
	RoomID    uuid.UUID `json:"room_id"`
	CreatedAt string    `json:"created_at"`
}

var db *sql.DB
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

const (
	tokenDuration     = 60 * time.Minute
	tokenRefreshLimit = 10 * time.Minute
)

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.ID = uuid.New()

	if _, err := db.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID.String(), user.Name, user.Email, string(hashedPassword)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var storedUser User
	if err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", user.Email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(tokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    storedUser.ID.String(),
		"name":  storedUser.Name,
		"email": storedUser.Email,
		"exp":   expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := extractTokenFromRequest(r)

	if _, err := db.Exec("INSERT INTO revoked_tokens (token) VALUES (?)", tokenString); err != nil {
		log.Println(err)
		http.Error(w, "Failed to revoke token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func tokenVerificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromRequest(r)

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if token is expired
		if exp, ok := (*claims)["exp"].(float64); !ok || time.Now().Unix() > int64(exp) {
			http.Error(w, "Expired token", http.StatusUnauthorized)
			return
		}

		// Check if token is close to its expiration time (less than 5 minutes left)
		if exp, ok := (*claims)["exp"].(float64); ok && time.Unix(int64(exp), 0).Sub(time.Now()) < tokenRefreshLimit {
			// Create a new token for the user with a new expiration time
			expirationTime := time.Now().Add(tokenDuration)
			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":    (*claims)["id"].(string),
				"name":  (*claims)["name"].(string),
				"email": (*claims)["email"].(string),
				"exp":   expirationTime.Unix(),
			})

			newTokenString, err := newToken.SignedString(jwtKey)
			if err != nil {
				http.Error(w, "Failed to generate new token", http.StatusInternalServerError)
				return
			}

			// Set new token as a cookie
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   newTokenString,
				Expires: expirationTime,
			})
		}

		next.ServeHTTP(w, r)
	})
}

func createChatRoomHandler(w http.ResponseWriter, r *http.Request) {
	var chatRoom ChatRoom
	if err := json.NewDecoder(r.Body).Decode(&chatRoom); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure chat room name is not blank
	if chatRoom.Name == "" {
		http.Error(w, "Chat room name cannot be blank", http.StatusBadRequest)
		return
	}

	chatRoom.ID = uuid.New()

	// Check if chat room with the same name exists
	var existingRoom ChatRoom
	err := db.QueryRow("SELECT id, name FROM chatrooms WHERE name = ?", chatRoom.Name).Scan(&existingRoom.ID, &existingRoom.Name)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Database error occurred", http.StatusInternalServerError)
		return
	}
	if existingRoom.Name != "" {
		http.Error(w, "Chat room with this name already exists", http.StatusConflict)
		return
	}

	_, err = db.Exec("INSERT INTO chatrooms (id, name) VALUES (?, ?)", chatRoom.ID.String(), chatRoom.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(chatRoom)
}

func chatRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM chatrooms")
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chatRooms)
}

func chatRoomHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var chatRoom ChatRoom
	if err := db.QueryRow("SELECT id, name FROM chatrooms WHERE id = ?", id).Scan(&chatRoom.ID, &chatRoom.Name); err != nil {
		http.Error(w, "Chatroom not found", http.StatusNotFound)
		return
	}

	rows, err := db.Query("SELECT name FROM users INNER JOIN room_user ON users.id = room_user.user_id WHERE room_user.room_id = ?", id)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]string, 0)
	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	chatRoom.Users = users

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	message.ID = uuid.New()

	_, err := db.Exec("INSERT INTO messages (id, text, sender_id, room_id, created_at) VALUES (?, ?, ?, ?, ?)", message.ID.String(), message.Text, message.SenderID.String(), message.RoomID.String(), time.Now())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func chatRoomMessagesHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	rows, err := db.Query("SELECT id, text, sender_id, room_id, created_at FROM messages WHERE room_id = ?", id)
	if err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	messages := make([]Message, 0)
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.Text, &message.SenderID, &message.RoomID, &message.CreatedAt); err != nil {
			http.Error(w, "Failed to retrieve message: "+err.Error(), http.StatusInternalServerError)
			return
		}

		messages = append(messages, message)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func main() {
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
	protected.HandleFunc("/chat/rooms", createChatRoomHandler).Methods("POST")
	protected.HandleFunc("/chat/rooms", chatRoomsHandler).Methods("GET")
	protected.HandleFunc("/chat/rooms/{id}", chatRoomHandler).Methods("GET")
	protected.HandleFunc("/chat/rooms/{id}/messages", sendMessageHandler).Methods("POST")
	protected.HandleFunc("/chat/rooms/{id}/messages", chatRoomMessagesHandler).Methods("GET")
	protected.HandleFunc("/chat/rooms/{roomID}/assign/{userID}", assignUserToChatRoom).Methods("POST")

	port := ":8000"

	fmt.Printf("Server started on port %s\n", port)
	http.ListenAndServe(port, router)
}
