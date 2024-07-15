package main

import (
	"catalk/config"
	"catalk/internal/database"
	"catalk/internal/server"
)

func main() {
	config := config.GetConfig()
	database := database.New(config.Database)
	server := server.NewServer(config, database)
	server.Start()
}
