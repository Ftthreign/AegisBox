package session

import (
	"context"
	"errors"
	"net/http"
	"time"
)

var (
	ErrNoSession       = errors.New("session not found")
	ErrSessionExpired  = errors.New("session expired")
)

type service struct {
	store  Store
	cookie Cookie
	cfg    Config
}

func New(store Store, cookie Cookie, cfg Config) Service {
	return &service{store: store, cookie: cookie, cfg: cfg}
}

// Create membuat session baru dan langsung set cookie-nya ke response writer.
func (s *service) Create(ctx context.Context, userID uint, w http.ResponseWriter) (string, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(time.Duration(s.cfg.SessionMaxAge()) * time.Second)

	token, err := s.store.CreateSession(ctx, userID, expiresAt)
	if err != nil {
		return "", err
	}

	s.cookie.Set(w, token, expiresAt)
	return token, nil
}

// Validate membaca cookie, cek token, cek expiry, lalu return userID.
func (s *service) Validate(ctx context.Context, r *http.Request) (uint, error) {
	token, err := s.cookie.Get(r)
	if err != nil {
		return 0, ErrNoSession
	}

	userID, expiresAt, err := s.store.GetSession(ctx, token)
	if err != nil {
		return 0, ErrNoSession
	}

	now := time.Now().UTC()
	if now.After(expiresAt) {
		_ = s.store.DeleteSession(ctx, token)
		return 0, ErrSessionExpired
	}

	return userID, nil
}

// Destroy menghancurkan session dari store dan cookie.
func (s *service) Destroy(ctx context.Context, r *http.Request, w http.ResponseWriter) error {
	token, err := s.cookie.Get(r)
	if err == nil {
		_ = s.store.DeleteSession(ctx, token)
	}

	s.cookie.Clear(w)
	return nil
}
