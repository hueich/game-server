package auth

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/hueich/game-server/db"
)

type UserManager struct {
	client *datastore.Client
}

func NewUserManager(ctx context.Context, projectID, credsFile string) (*UserManager, error) {
	c, err := db.NewDatastoreClient(ctx, projectID, credsFile)
	if err != nil {
		return nil, err
	}
	return &UserManager{client: c}, nil
}
