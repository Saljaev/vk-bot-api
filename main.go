package main

import (
	"go/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := handlers.GetLongPollServer(); err != nil {
		panic(err)
	}

	go handlers.LongPollHandler()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
