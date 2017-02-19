package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	bapi "github.com/hueich/blokus-web-api"
)

var (
	port = flag.Int("port", 9090, "Port to start the server on")
)

func main() {
	log.Println("DATASTORE_PROJECT_ID:", os.Getenv("DATASTORE_PROJECT_ID"))
	log.Println("GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	r := mux.NewRouter()
	br := r.PathPrefix("/blokus").Subrouter()
	s, err := bapi.NewService(br.PathPrefix("/api").Subrouter())
	if err != nil {
		log.Fatalf("Could not create service: %v\n", err)
	}
	defer s.Close()
	if err := s.InitClient(context.Background(), "", ""); err != nil {
		log.Fatalf("Could not initialize client: %v\n", err)
	}

	http.Handle("/", r)
	log.Println("Starting to listen on port", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
