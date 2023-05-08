package main

import (
	"go/handlers"
	"net/http"
)

func main() {
	if err := handlers.GetLongPollServer(); err != nil {
		panic(err)
	}

	go handlers.LongPollHandler()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
