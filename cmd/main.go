package main

import (
	"fmt"
	"log"

	"github.com/nicolekistler/ws-lobby/internal/server"
)

func main() {
	fmt.Println("Main")
	serv := server.NewServer()

	if err := serv.Init(); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
