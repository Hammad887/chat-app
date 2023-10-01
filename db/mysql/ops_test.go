package mysql

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Hammad887/chat-app/db"
	domain "github.com/Hammad887/chat-app/models"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

func Test_client_LogoutUser(t *testing.T) {
	os.Setenv("MYSQL_DB_HOSTS", "mysql_db")

	options := db.Option{
		TestMode: false, // Set the appropriate value for TestMode here
	}
	client, _ := NewClient(options)

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
			got, err := client.LogoutUser(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LogoutUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.LogoutUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_RegisterUser(t *testing.T) {
	os.Setenv("MYSQL_DB_HOSTS", "mysql_db")

	options := db.Option{
		TestMode: false, // Set the appropriate value for TestMode here
	}
	client, _ := NewClient(options)

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
		// test cases
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
			got, err := client.RegisterUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.RegisterUser() = %v, want %v", got, tt.want)
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
