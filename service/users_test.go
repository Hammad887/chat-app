package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Hammad887/chat-app/db"
	"github.com/Hammad887/chat-app/db/mysql"
	domain "github.com/Hammad887/chat-app/models"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

func Test_service_LogoutUser(t *testing.T) {
	os.Setenv("MYSQL_DB_HOSTS", "mysql_db")

	options := db.Option{
		TestMode: false, // Set the appropriate value for TestMode here
	}
	client, _ := mysql.NewClient(options)
	s := NewService(&client)

	user := getUser()

	type args struct {
		ctx   context.Context
		token string
	}

	storedUUID := user.ID

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// test cases.
		{
			name:    "success case of user logout",
			args:    args{ctx: context.Background(), token: storedUUID},
			wantErr: false,
			want:    true,
		},
		{
			name:    "failure case of user logout",
			args:    args{ctx: context.Background(), token: storedUUID},
			wantErr: true,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.LogoutUser(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogoutUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("client.LogoutUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_RegisterUser(t *testing.T) {
	os.Setenv("MYSQL_DB_HOSTS", "mysql_db")

	options := db.Option{
		TestMode: false, // Set the appropriate value for TestMode here
	}
	client, _ := mysql.NewClient(options)
	s := NewService(&client)

	storedUser := getUser()
	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "success case of user registration",
			args:    args{ctx: context.Background(), user: storedUser},
			wantErr: false,
			want:    true,
		},
		{
			name:    "failure case of user registration",
			args:    args{ctx: context.Background(), user: storedUser},
			wantErr: true,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RegisterUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.RegisterUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getUser() *domain.User {
	return &domain.User{
		ID:       uuid.New().String(),
		Name:     randstr.String(12),
		Email:    fmt.Sprintf("%s@case.com", randstr.String(8)),
		Password: randstr.String(12),
	}
}
