package main

import (
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

	server := NewAPIServer(":3000", store)
	server.Run()
}
