package main

import (
	"fmt"
	"final-project-backend/db"
	"final-project-backend/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("failed to connect to DB")
	}

	server.Init()
}
