package auth

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/hueich/game-server/db"
)

type userInfo struct {
	Username string
	PasswordHash []byte
	Email string
}

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

func (m *UserManager) AddUser(username string, password []byte, email string) error {
	// TODO: Check for existing user in db

	// TODO: Hash password with bcrypt

	// TODO: Add user to db

	return nil
}
