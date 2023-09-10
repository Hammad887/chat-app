package mysql

import (
	"context"
	"database/sql"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, fmt.Errorf("failed to hash password: %w", err)
	}

	user.ID = generateUserID()

	if _, err := c.dbc.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, string(hashedPassword)); err != nil {
		return false, fmt.Errorf("failed to execute database insert: %w", err)
	}

	// createChatRoom()
	return true, nil
}

func (c *client) SendMessage(ctx context.Context, id string, message *domain.Message) error {
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
	rows, err := c.dbc.Query("SELECT id, name FROM chatrooms")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return scanChatRooms(rows)
}

func scanChatRooms(rows *sql.Rows) ([]*domain.ChatRoom, error) {
	chatRooms := make([]*domain.ChatRoom, 0)

	for rows.Next() {
		var chatRoom domain.ChatRoom
		if err := rows.Scan(&chatRoom.ID, &chatRoom.Name); err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		chatRooms = append(chatRooms, &chatRoom)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chatRooms, nil
}

func (c *client) GetChatroom(ctx context.Context, id string) (*domain.ChatRoom, error) {
	var chatRoom domain.ChatRoom
	if err := c.dbc.QueryRow("SELECT id, name FROM chatrooms WHERE id = ?", id).Scan(&chatRoom.ID, &chatRoom.Name); err != nil {
		return nil, fmt.Errorf("failed to scan row data: %w", err)
	}

	rows, err := c.dbc.Query("SELECT name FROM users INNER JOIN room_user ON users.id = room_user.user_id WHERE room_user.room_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	users := make([]string, 0)
	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	chatRoom.Users = users

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

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

const (
	tokenDuration = 60 * time.Minute
)

func (c *client) LoginUser(ctx context.Context, email string, password string) (string, error) {
	var storedUser domain.User
	if err := c.dbc.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password); err != nil {
		return "", err
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
