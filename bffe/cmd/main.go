package main

import (
	"log"

	"github.com/o-rensa/jalikod-backend/bffe/cmd/api"
)

func main() {
	server := api.NewAPIServer(3000)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
