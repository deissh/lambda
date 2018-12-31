package main

import (
	"github.com/deissh/lambda/pkg/server"
	"log"
	"os"
)

func main() {
	if os.Getenv("DEBUG") == "true" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	log.Println("starting server")

	log.Fatal(server.Run("0.0.0.0", "8080"))
}
