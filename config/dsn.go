package config

import (
	"fmt"
	"os"
)

var (
	dbUser     = os.Getenv("PSQL_DB_USER")
	dbPassword = os.Getenv("PSQL_DB_PASSWORD")
	dbName     = os.Getenv("PSQL_DB_NAME")
	dbHost     = os.Getenv("PSQL_DB_HOST")
	dbPort     = os.Getenv("PSQL_DB_PORT")
	dbSSLMode  = os.Getenv("PSQL_DB_SSLMODE")

	mariaDbUser     = os.Getenv("MARIA_DB_USER")
	mariaDbPassword = os.Getenv("MARIA_DB_PASSWORD")
	mariaDbName     = os.Getenv("MARIA_DB_NAME")
	mariaDbHost     = os.Getenv("MARIA_DB_HOST")
	mariaDbPort     = os.Getenv("MARIA_DB_PORT")
)

// NewPostgresDsnFromEnv provides the dsn connection string for database
func NewPostgresDsnFromEnv() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)
}

// NewPostgresDsnFromEnv provides the dsn connection string for database
func NewMariaDsnFromEnv() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mariaDbUser, mariaDbPassword, mariaDbName, mariaDbHost, mariaDbPort)
}
