package main

import (
	"log"
	"net/http"

	bs "github.com/hueich/blokus-service"
)

func main() {
	r := bs.NewRouter("/blokus/")
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
