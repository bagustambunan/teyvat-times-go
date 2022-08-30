package main

import (
	"final-project-backend/db"
	"final-project-backend/server"
	"fmt"
)

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("failed to connect to DB")
	}

	server.Init()
}
