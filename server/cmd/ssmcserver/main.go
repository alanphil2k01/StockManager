package main

import (
	"log"

	"github.com/alanphil2k01/SSMC/pkg/server"
)

func main() {
	errC, err := server.RunServer()
	if err != nil {
		log.Fatalf("Couldn not run the server: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running the server: %s", err)
	}
}
