package main

import (
	"log"

	"github.com/harsh/project/cmd/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.Server()

}
