package session

import (
	"context"
	"net/http"
	"time"
)

// Store (infrastructure) harus implement interface ini.
// Dia bertugas nyimpen token, load token, destroy token.
type Store interface {
	CreateSession(ctx context.Context, userID uint, expiresAt time.Time) (string, error)
	GetSession(ctx context.Context, token string) (userID uint, expiresAt time.Time, err error)
	DeleteSession(ctx context.Context, token string) error
}

// Cookie (infrastructure) bertugas atur cookie HTTP.
type Cookie interface {
	Set(w http.ResponseWriter, token string, expiresAt time.Time)
	Clear(w http.ResponseWriter)
	Get(r *http.Request) (string, error)
}

type Config interface {
	SessionMaxAge() int            // dalam detik
	SessionIdleTimeout() int       // detik
}

type Service interface {
	Create(ctx context.Context, userID uint, w http.ResponseWriter) (string, error)
	Validate(ctx context.Context, r *http.Request) (uint, error)
	Destroy(ctx context.Context, r *http.Request, w http.ResponseWriter) error
}
