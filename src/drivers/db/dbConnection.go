package db

import (
	"context"
	"database/sql"
	"eniqilo-store/src/config"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {

	strConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME, config.DB_PARAMS)

	// Define connection pool parameters (adjust as needed)
	maxOpenConns := 20 // Maximum number of open connections in the pool
	maxIdleConns := 10 // Maximum number of idle connections in the pool

	db, err := sql.Open("postgres", strConnection)
	if err != nil {
		return nil, err
	}

	// Create connection pool
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	// Test connection using PingContext
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		db.Close() // Close the connection pool on error
		return nil, err
	}

	return db, nil
}
