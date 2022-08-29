package main

import (
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/server"
	"fmt"
)

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("failed to connect to DB")
	}

	server.Init()
}
