package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/datastore"
)

const (
	projectID = "axial-analyzer-117623"
)

type Coord struct {
	X, Y int
}

type Piece struct {
	Name      string
	Coords    []Coord
	CreatedAt time.Time
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.GET("/connect", connect)

	// Listen and serve on port
	router.Run(":" + port)
}

func connect(c *gin.Context) {
	// client, err := google.DefaultClient(context.Background(), datastore.ScopeDatastore)
	//    if err != nil {
	//            log.Fatalf("Unable to get default client: %v", err)
	//    }
	// service, err := storage.New(client)
	//    if err != nil {
	//            log.Fatalf("Unable to create storage service: %v", err)
	//    }

	jsonKey := os.Getenv("KEY_JSON")
	// jsonKey, err := ioutil.ReadFile("/path/to/json/keyfile.json")
	if jsonKey == "" {
		log.Fatal("KEY_JSON is empty!")
	}
	conf, err := google.JWTConfigFromJSON(
		[]byte(jsonKey),
		datastore.ScopeDatastore,
		datastore.ScopeUserEmail,
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID, cloud.WithTokenSource(conf.TokenSource(ctx)))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected!")

	key := datastore.NewIncompleteKey(ctx, "Piece", nil)
	tx, err := client.NewTransaction(ctx)
	if err != nil {
		log.Fatal(err)
	}

	piece := &Piece{Name: "somepiece", CreatedAt: time.Now()}
	if _, err := tx.Put(key, &piece); err != nil {
		log.Fatal(err)
	}

	// Attempt to commit the transaction. If there's a conflict, try again.
	if _, err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	log.Print("Committed!")

	c.String(http.StatusOK, "Updated database!")
}
