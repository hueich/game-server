package auth

import (
	"cloud.google.com/go/datastore"
)

type UserManager struct {
	client *datastore.Client
}

func NewUserManager() (*UserManager, error) {
	return &UserManager{}
}
