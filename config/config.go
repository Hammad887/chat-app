package config

import (
	"log"

	"github.com/spf13/viper"
)

// keys for database configuration
const (
	MySQLDBHost     = "MYSQL_DB_HOST"
	MySQLDBUsername = "MYSQL_DB_USERNAME"
	MySQLDBPassword = "MYSQL_DB_PASSWORD"
	MySQLDBName     = "MYSQL_DB_NAME"
)

func init() {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env") // Assuming you're using .env in the root
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	log.Println(viper.AllSettings())

	if err != nil {
		log.Printf("Error reading config file: %s\n", err)
	}

	// env var for db
	_ = viper.BindEnv(MySQLDBHost, "MYSQL_DB_HOSTS")
	_ = viper.BindEnv(MySQLDBUsername, "MYSQL_DB_USERNAME")
	_ = viper.BindEnv(MySQLDBPassword, "MYSQL_DB_PASSWORD")

	viper.SetDefault(MySQLDBName, "userdb")
	viper.SetDefault(MySQLDBHost, "127.0.0.1:3306")
}
