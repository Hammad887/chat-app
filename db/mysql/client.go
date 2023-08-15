package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	wraperrors "github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Hammad887/chat-app/config"
	"github.com/Hammad887/chat-app/db"
)

type client struct {
	dbc *sql.DB
}

// NewClient creates a configured database client.
func NewClient(option db.Option) (db.DataStore, error) {
	ctx := context.Background()

	username := viper.GetString(config.MySQLDBUsername)
	password := viper.GetString(config.MySQLDBPassword)
	hostname := viper.GetString(config.MySQLDBHost)
	dbName := viper.GetString(config.MySQLDBName)
	// username := "hammad"
	// password := "Hammad_887"
	// hostname := "127.0.0.1:3306"
	// dbName := "chatapp"

	fmt.Println(dsn(username, password, hostname, dbName))

	// dbClient, err := sql.Open("mysql", "abc:password@tcp(db:3306)/test_database")

	dbClient, err := sql.Open("mysql", "hammad:Hammad_887@tcp(localhost:3306)/chatapp")

	// dbClient, err := sql.Open("mysql", "newuser:password@tcp(db:3306)/chatapp")
	if err != nil {
		return nil, wraperrors.Wrap(err, "unable to connect to database")
	}

	if err := dbClient.PingContext(ctx); err != nil {
		return nil, wraperrors.Wrap(err, "unable to ping database")
	}

	return &client{dbc: dbClient}, nil
}

func dsn(username string, password string, hostname string, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
