package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	bs "github.com/hueich/blokus-service"
)

var (
	port = flag.Int("port", 9090, "Port to start the server on")
)

func main() {
	log.Println("DATASTORE_PROJECT_ID:", os.Getenv("DATASTORE_PROJECT_ID"))
	log.Println("GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	s := bs.New("/blokus/")
	if err := s.InitClient(context.Background(), "", ""); err != nil {
		log.Fatalf("Could not create client: %v\n", err)
	}
	defer s.Close()

	http.Handle("/", s.Router())
	log.Println("Starting to listen on port", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
