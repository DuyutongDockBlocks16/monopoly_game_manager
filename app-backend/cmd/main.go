package main

import (
	"fmt"
	"log"
	"app-backend/internal/server"
	"app-backend/database"
)

func init() {
	database.NewDB()
}

func main() {
	port := "3000"
	fmt.Println("Server is running on port", port)

	err := server.Start(":" + port)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
