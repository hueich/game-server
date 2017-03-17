package auth

import (
	"errors"

	"github.com/gorilla/sessions"
)

const (
	DefaultSessionName = "gs"
)

type SessionManager struct {
	Name string

	store sessions.Store
}

type SessionManagerOptions struct {
	SessionName string
}

func NewSessionManager(store sessions.Store, opt SessionManagerOptions) (*SessionManager, error) {
	if store == nil {
		return nil, errors.New("session store cannot be nil")
	}
	m := &SessionManager{
		Name:  DefaultSessionName,
		store: store,
	}
	if opt.SessionName != "" {
		m.Name = opt.SessionName
	}
	return m, nil
}
