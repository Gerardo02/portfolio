package main

import (
	"log"

	"github.com/gerardo02/porfolio/cmd/api"
)

func main() {
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
