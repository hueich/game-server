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
	bapp "github.com/hueich/blokus/web/app"
	bapi "github.com/hueich/blokus/web/rest"
	"github.com/hueich/game-server/db"
)

var (
	port = flag.Int("port", 9090, "Port to start the server on")
)

func newSessionStore() (sessions.Store, error) {
	// DEBUG
	log.Println("SERVER_SESSION_HASH_KEY:", os.Getenv("SERVER_SESSION_HASH_KEY"))
	log.Println("SERVER_SESSION_ENCRYPT_KEY:", os.Getenv("SERVER_SESSION_ENCRYPT_KEY"))

	hkey, err := ioutil.ReadFile(os.Getenv("SERVER_SESSION_HASH_KEY"))
	if err != nil {
		return nil, fmt.Errorf("Could not read session hash key file: %v", err)
	}
	if len(hkey) == 0 {
		return nil, fmt.Errorf("Hash key content is empty")
	}

	ekey, err := ioutil.ReadFile(os.Getenv("SERVER_SESSION_ENCRYPT_KEY"))
	if err != nil {
		return nil, fmt.Errorf("Could not read session encryption key file: %v", err)
	}
	if len(ekey) == 0 {
		return nil, fmt.Errorf("Encryption key content is empty")
	}

	store := sessions.NewCookieStore(hkey, ekey)
	store.Options.MaxAge = 1 * 24 * 60 * 60 // 1 day
	store.Options.HttpOnly = true
	return store, nil
}

func run() error {
	store, err := newSessionStore()
	if err != nil {
		return fmt.Errorf("Could not initialize session store: %v", err)
	}

	// DEBUG
	log.Println("BLOKUS_PROJECT_ID:", os.Getenv("BLOKUS_PROJECT_ID"))
	log.Println("BLOKUS_GAPP_CREDS:", os.Getenv("BLOKUS_GAPP_CREDS"))

	r := mux.NewRouter()
	br := r.PathPrefix("/blokus").Subrouter()

	dsClient, err := db.NewDatastoreClient(context.Background(), os.Getenv("BLOKUS_PROJECT_ID"), os.Getenv("BLOKUS_GAPP_CREDS"))
	if err != nil {
		return err
	}
	defer dsClient.Close()

	sAPI, err := bapi.NewService(bapi.Options{
		Router: br.PathPrefix("/api").Subrouter(),
		Client: dsClient,
	})
	if err != nil {
		return fmt.Errorf("Could not create API service: %v", err)
	}
	defer sAPI.Close()

	_, err = bapp.NewService(&bapp.Options{
		Router: br,
		ApiURL: "/blokus/api",
		Store:  store,
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

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
