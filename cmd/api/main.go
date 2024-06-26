package main

import (
	"catalk/internal/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
