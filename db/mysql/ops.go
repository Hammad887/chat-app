package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	domain "github.com/Hammad887/chat-app/models"
)

// generateUserID creates a new UUID for user ID.
func generateUserID() string {
	return uuid.New().String()
}

func (c *client) RegisterUser(ctx context.Context, user *domain.User) (bool, error) {
	user.ID = generateUserID()

	if _, err := c.dbc.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Password); err != nil {
		return false, fmt.Errorf("failed to execute database insert: %w", err)
	}

	// createChatRoom()
	return true, nil
}

func (c *client) SaveMessage(ctx context.Context, id string, message *domain.Message) error {
	message.ID = uuid.New().String()

	// c.assignUserToChatRoom()

	_, err := c.dbc.Exec("INSERT INTO messages (id, text, sender_id, room_id, created_at) VALUES (?, ?, ?, ?, ?)", message.ID, message.Text, message.SenderID, message.RoomID, time.Now())
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to execute database update: %w", err)
	}

	return nil
}

func (c *client) ListChatRoom(ctx context.Context) ([]*domain.ChatRoom, error) {
	rows, err := c.dbc.Query("SELECT c.id, c.name, ru.user_id FROM chatrooms c JOIN room_user ru ON c.id = ru.room_id ")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return scanChatRooms(rows)
}

func (c *client) GetChatroom(ctx context.Context, id string) (*domain.ChatRoom, error) {
	var chatRoom domain.ChatRoom

	rows, err := c.dbc.Query("SELECT c.id, c.name, ru.user_id FROM chatrooms c JOIN room_user ru ON c.id = ru.room_id WHERE c.id = ?", id)
	if err != nil {
		return nil, err
	}
	userIDs := make(map[string]struct{})

	// Iterate through the rows and scan data into the struct and user IDs
	for rows.Next() {
		var userID string
		if err := rows.Scan(&chatRoom.ID, &chatRoom.Name, &userID); err != nil {
			log.Fatal(err)
		}
		chatRoom.Users = append(chatRoom.Users, userID)
		userIDs[userID] = struct{}{}
	}

	return &chatRoom, nil
}

func (c *client) GetChatroomMessages(ctx context.Context, id string) ([]*domain.Message, error) {
	rows, err := c.dbc.Query("SELECT id, text, sender_id, room_id, created_at FROM messages WHERE room_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	messages := make([]*domain.Message, 0)
	for rows.Next() {
		var message domain.Message
		if err := rows.Scan(&message.ID, &message.Text, &message.SenderID, &message.RoomID, &message.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan message row: %w", err)
		}

		messages = append(messages, &message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *client) createChatRoom(name string) error {

	// Ensure chat room name is not blank
	if name == "" {
		return errors.New("chat room name cannot be blank")
	}

	ID := uuid.New().String()

	// Check if chat room with the same name exists
	var existingRoom domain.ChatRoom
	err := c.dbc.QueryRow("SELECT id, name FROM chatrooms WHERE name = ?", name).Scan(&existingRoom.ID, &existingRoom.Name)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("database error occurred")
	}

	if existingRoom.Name != "" {
		return errors.New("chat room with this name already exists")
	}

	_, err = c.dbc.Exec("INSERT INTO chatrooms (id, name) VALUES (?, ?)", ID, name)
	if err != nil {
		return errors.New("could not create new chatroom")
	}

	return nil
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

const (
	tokenDuration = 60 * time.Minute
)

func (c *client) LoginUser(ctx context.Context, email string, password string) (string, error) {
	var storedUser domain.User
	if err := c.dbc.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password); err != nil {
		return "", fmt.Errorf("failed to scan row: %w", err)
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("failed to scan row: %w", err)
	}

	expirationTime := time.Now().Add(tokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    storedUser.ID,
		"name":  storedUser.Name,
		"email": storedUser.Email,
		"exp":   expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to get signed string from JWT token: %w", err)
	}

	return tokenString, nil
}

func (c *client) LogoutUser(ctx context.Context, token string) (bool, error) {
	if _, err := c.dbc.Exec("INSERT INTO revoked_tokens (token) VALUES (?)", token); err != nil {
		log.Println(err)
		return false, fmt.Errorf("failed to scan row: %w", err)
	}

	return true, nil
}

func scanChatRooms(rows *sql.Rows) ([]*domain.ChatRoom, error) {
	chatRoomMap := make(map[string]*domain.ChatRoom)

	for rows.Next() {
		var chatRoom domain.ChatRoom
		var userId string

		if err := rows.Scan(&chatRoom.ID, &chatRoom.Name, &userId); err != nil {
			return nil, fmt.Errorf("failed to scan chatroom row: %w", err)
		}

		if existingChatRoom, exists := chatRoomMap[chatRoom.ID]; exists {
			existingChatRoom.Users = append(existingChatRoom.Users, userId)
		} else {
			chatRoom.Users = append(chatRoom.Users, userId)
			chatRoomMap[chatRoom.ID] = &chatRoom
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	chatRooms := make([]*domain.ChatRoom, 0, len(chatRoomMap))
	for _, chatRoom := range chatRoomMap {
		chatRooms = append(chatRooms, chatRoom)
	}

	return chatRooms, nil
}
