package mysql

import (
	"context"
	"database/sql"
	"fmt"

	// The following import is required to register the MySQL driver with the database/sql package.
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"github.com/Hammad887/chat-app/config"
	"github.com/Hammad887/chat-app/db"
)

type client struct {
	dbc *sql.DB
}

// NewClient creates a configured database client.
func NewClient(_ db.Option) (db.DataStore, error) {
	ctx := context.Background()

	username := viper.GetString(config.MySQLDBUsername)
	password := viper.GetString(config.MySQLDBPassword)
	hostname := viper.GetString(config.MySQLDBHost)
	dbName := viper.GetString(config.MySQLDBName)

	dbClient, err := sql.Open("mysql", dsn(username, password, hostname, dbName))

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := dbClient.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &client{dbc: dbClient}, nil
}

func dsn(username string, password string, hostname string, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
