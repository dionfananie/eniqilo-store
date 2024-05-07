package database

import (
	"database/sql"
	"fmt"
	"log"
	"web/go-cat-match/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error

	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME, config.DB_PARAMS)

	DB, err = sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	println("database connected!")
}
