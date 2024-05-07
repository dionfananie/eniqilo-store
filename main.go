package main

import (
	"eniqilo-store/src/drivers/db"
	"eniqilo-store/src/http"
	"fmt"
)

func main() {
	dbConnection, err := db.CreateConnection()
	if err != nil {
		fmt.Println("Error creating database connection:", err)
		return
	}

	defer func() {
		if err := dbConnection.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	h := http.New(&http.Http{
		DB: dbConnection,
	})
	h.Launch()
}
