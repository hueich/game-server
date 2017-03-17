package auth

import (
	"errors"

	"github.com/gorilla/sessions"
)

type SessionManager struct {
	store sessions.Store
}

func NewSessionManager(store sessions.Store) (*SessionManager, error) {
	if store == nil {
		return nil, errors.New("session store cannot be nil")
	}
	m := &SessionManager{
		store: store,
	}
	return m, nil
}
