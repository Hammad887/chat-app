package service

import (
	"github.com/yourapp/models"
	"github.com/yourapp/db"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.ID = uuid.New()

	_, err = db.DB.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID.String(), user.Name, user.Email, string(hashedPassword))

	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return &user, err
}
