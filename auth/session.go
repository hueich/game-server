package auth

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	defaultSessionName = "ss"
)

var (
	ErrNotLoggedIn = errors.New("session: user not logged in")
)

type SessionManager struct {
	Name string

	store sessions.Store
}

type SessionManagerOptions struct {
	SessionName string
}

func NewSessionManager(store sessions.Store, opt *SessionManagerOptions) (*SessionManager, error) {
	if store == nil {
		return nil, errors.New("session: store cannot be nil")
	}
	m := &SessionManager{
		Name:  defaultSessionName,
		store: store,
	}
	if opt != nil {
		if opt.SessionName != "" {
			m.Name = opt.SessionName
		}
	}
	return m, nil
}

func (m *SessionManager) GetUsername(r *http.Request) (string, error) {
	s, err := m.store.Get(r, m.Name)
	if err != nil {
		return "", err
	}
	if s.IsNew {
		return "", ErrNotLoggedIn
	}
	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		return "", errors.New("session: username is empty")
	}
	return username, nil
}
