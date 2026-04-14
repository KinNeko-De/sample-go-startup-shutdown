package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	dbHost           = "localhost"
	dbUser           = "sampleuser"
	dbPassword       = "samplepassword"
	dbName           = "sampledatabase"
	dbPort           = "23100"
	dbSSLMode        = "disable"
	dbConnectTimeout = "5"
	maxRetries       = 5
	retryDelay       = 1 * time.Second
)

var Database *sql.DB

func ConnectToPostgres(logger *slog.Logger, connected chan<- struct{}) {
	logger.Info("connecting to postgress...")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s connect_timeout=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbConnectTimeout,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("Failed to open database connection", "error", err)
		return
	}

	err = pingDatabase(db, logger)
	if err != nil {
		logger.Error("Failed to ping database after multiple attempts", "error", err)
		return
	}

	Database = db
	logger.Info("connected to postgress")

	close(connected)
}

func pingDatabase(db *sql.DB, logger *slog.Logger) error {
	var err error
	for i := range maxRetries {
		err = db.Ping()
		if err == nil {
			return nil
		}
		logger.Info("ping to database failed", "error", err, "attempt", i+1, "max_retries", maxRetries)
		time.Sleep(retryDelay)
	}

	return err
}

func DisconnectFromPostgres(logger *slog.Logger) {
	logger.Info("disconnecting from postgress...")
	if Database != nil {
		err := Database.Close()
		if err != nil {
			logger.Error("Failed to close database connection", "error", err)
		}
		logger.Info("disconnected from postgress")
	} else {
		logger.Info("database connection was not established yet")
	}
}
