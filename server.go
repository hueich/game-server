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
	bui "github.com/hueich/blokus-web-ui"
)

var (
	port = flag.Int("port", 9090, "Port to start the server on")
)

func main() {
	log.Println("DATASTORE_PROJECT_ID:", os.Getenv("DATASTORE_PROJECT_ID"))
	log.Println("GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	r := mux.NewRouter()
	br := r.PathPrefix("/blokus").Subrouter()

	sAPI, err := bapi.NewService(br.PathPrefix("/api").Subrouter())
	if err != nil {
		log.Fatalf("Could not create API service: %v\n", err)
	}
	defer sAPI.Close()
	if err := sAPI.InitDBClient(context.Background(), "", ""); err != nil {
		log.Fatalf("Could not initialize client: %v\n", err)
	}

	_, err = bui.NewService(br, "/blokus/api", "")
	if err != nil {
		log.Fatalf("Could not create UI service: %v\n", err)
	}

	http.Handle("/", r)
	log.Println("Starting to listen on port", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
