package main

import (
	"log"
	"os"

	"github.com/gerardo02/porfolio/internal/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Printf("Server error: %v", err)
		os.Exit(1)
	}
}
