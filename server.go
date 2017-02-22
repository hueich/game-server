package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	bapi "github.com/hueich/blokus-web-api"
	bapp "github.com/hueich/blokus-web-app"
)

var (
	port = flag.Int("port", 9090, "Port to start the server on")

	sessionStore sessions.Store
)

func initSessionStore() error {
	// DEBUG
	log.Println("SERVER_SESSION_HASH_KEY:", os.Getenv("SERVER_SESSION_HASH_KEY"))
	log.Println("SERVER_SESSION_ENCRYPT_KEY:", os.Getenv("SERVER_SESSION_ENCRYPT_KEY"))

	hkey, err := ioutil.ReadFile(os.Getenv("SERVER_SESSION_HASH_KEY"))
	if err != nil {
		return fmt.Errorf("Could not read session hash key file: %v", err)
	}
	if len(hkey) == 0 {
		return fmt.Errorf("Hash key content is empty")
	}

	ekey, err := ioutil.ReadFile(os.Getenv("SERVER_SESSION_ENCRYPT_KEY"))
	if err != nil {
		return fmt.Errorf("Could not read session encryption key file: %v", err)
	}
	if len(ekey) == 0 {
		return fmt.Errorf("Encryption key content is empty")
	}

	store := sessions.NewCookieStore(hkey, ekey)
	store.Options.MaxAge = 1 * 24 * 60 * 60 // 1 day
	store.Options.HttpOnly = true
	sessionStore = store
	return nil
}

func run() error {
	if err := initSessionStore(); err != nil {
		return fmt.Errorf("Could not initialize session store: %v", err)
	}

	r := mux.NewRouter()
	br := r.PathPrefix("/blokus").Subrouter()

	sAPI, err := bapi.NewService(br.PathPrefix("/api").Subrouter())
	if err != nil {
		return fmt.Errorf("Could not create API service: %v", err)
	}
	defer sAPI.Close()
	if err := sAPI.InitDBClient(context.Background(), "", ""); err != nil {
		return fmt.Errorf("Could not initialize client: %v", err)
	}

	_, err = bapp.NewService(&bapp.Options{
		Router: br,
		ApiURL: "/blokus/api",
		Store:  sessionStore,
	})
	if err != nil {
		return fmt.Errorf("Could not create app service: %v", err)
	}

	http.Handle("/", r)
	log.Println("Starting to listen on port", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		return fmt.Errorf("Could not start server: %v", err)
	}
	return nil
}

func main() {
	flag.Parse()

	// DEBUG
	log.Println("DATASTORE_PROJECT_ID:", os.Getenv("DATASTORE_PROJECT_ID"))
	log.Println("GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
