package main

import (
	"log"
	"net"
	"net/http"
	"os"

	a "github.com/christinavanmoof/generatePasswords/api"
	b "github.com/christinavanmoof/generatePasswords/passwords"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Printf("Could not listen on port 9000: %v", err)
		os.Exit(1)
	}
	log.Printf("Listening on port 9000")

	// Interface that is generating the passwords
	c := &b.GeneratePasswords{}

	// Package that is handling the requests
	handler := &a.API{
		Gen: c,
	}

	// Serve the requests
	if err := http.Serve(lis, handler); err != nil {
		log.Printf("Could not serve: %v", err)
		os.Exit(1)
	}
}
