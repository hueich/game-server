package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hueich/game-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/hueich/game-server/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/hueich/game-server/Godeps/_workspace/src/golang.org/x/oauth2/google"
	"github.com/hueich/game-server/Godeps/_workspace/src/google.golang.org/cloud"
	"github.com/hueich/game-server/Godeps/_workspace/src/google.golang.org/cloud/datastore"
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
	Corner    Coord
	M         map[Coord]int
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

func getClient(ctx context.Context) (*datastore.Client, error) {
	jsonKey := os.Getenv("KEY_JSON")
	if jsonKey == "" {
		return nil, fmt.Errorf("KEY_JSON is empty!")
	}
	conf, err := google.JWTConfigFromJSON(
		[]byte(jsonKey),
		datastore.ScopeDatastore,
		datastore.ScopeUserEmail,
	)
	if err != nil {
		return nil, err
	}
	client, err := datastore.NewClient(ctx, projectID, cloud.WithTokenSource(conf.TokenSource(ctx)))
	if err != nil {
		return nil, err
	}
	log.Print("Got client!")
	return client, nil
}

func connect(c *gin.Context) {
	ctx := context.Background()
	client, err := getClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	key := datastore.NewIncompleteKey(ctx, "Piece", nil)
	tx, err := client.NewTransaction(ctx)
	if err != nil {
		log.Fatal(err)
	}

	piece := &Piece{Name: "pieceWithCoords", Corner: Coord{3, 4}, M: make(map[Coord]int), CreatedAt: time.Now()}
	piece.Coords = append(piece.Coords, Coord{1, 2})
	piece.Coords = append(piece.Coords, Coord{5, 6})
	piece.Coords = append(piece.Coords, Coord{100, 300})
	piece.M[Coord{9, 9}] = 777
	if _, err := tx.Put(key, piece); err != nil {
		log.Fatal(err)
	}

	if _, err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	log.Print("Committed!")

	c.String(http.StatusOK, "Updated database!")
}
