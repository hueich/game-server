package auth

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/hueich/game-server/db"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLength = 6
)

type userInfo struct {
	Username     string
	PasswordHash []byte `datastore:",noindex"`
	Email        string
	TimeCreated  time.Time
	TimeLastSeen time.Time
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

func (m *UserManager) AddUser(ctx context.Context, username string, password []byte, email string) (int64, error) {
	if username == "" {
		return 0, fmt.Errorf("username cannot be empty")
	}
	if len(password) < MinPasswordLength {
		return 0, fmt.Errorf("password is too short")
	}
	if email != "" {
		addr, err := mail.ParseAddress(email)
		if err != nil {
			return 0, fmt.Errorf("email is invalid")
		}
		email = addr.Address
	}

	var pk *datastore.PendingKey
	commit, err := m.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		q := datastore.NewQuery("User").Transaction(tx).Limit(1).Filter("Username =", username)
		n, err := m.client.Count(ctx, q)
		if err != nil {
			return err
		}
		if n > 0 {
			return fmt.Errorf("user already exists")
		}

		hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to generate hash: %v", err)
		}

		u := &userInfo{
			Username:     username,
			PasswordHash: hash,
			Email:        email,
			TimeCreated:  time.Now(),
			TimeLastSeen: time.Now(),
		}

		pk, err = tx.Put(datastore.IncompleteKey("User", nil), u)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to add new user: %v", err)
	}
	key := commit.Key(pk)
	return key.ID, nil
}
