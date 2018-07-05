package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Duration(5 * time.Millisecond),
	}

	_, err := client.Get("https://golangcode.com")
	if err != nil {
		log.Fatal(err)
	}
}
