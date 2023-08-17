package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	store, err := NewPostgresStore(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	port = fmt.Sprintf(":%s", port)
	server := NewAPIServer(port, store)
	server.Run()
}
