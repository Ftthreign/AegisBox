package sessionInfra

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type SessionData struct {
	UserID    uint
	ExpiresAt time.Time
}

type MemoryStore struct {
	mu 	     sync.RWMutex
	sessions map[string]SessionData
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		sessions: make(map[string]SessionData),
	}
}

func (m *MemoryStore) CreateSession(
	ctx context.Context, 
	userID uint, 
	expiresAt time.Time) (string, error) {
	token := uuid.New().String()

	m.mu.Lock()
	m.sessions[token] = SessionData{
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
	m.mu.Unlock()

	return token, nil
}

func (m *MemoryStore) GetSession(
	ctx context.Context, 
	token string) (uint, time.Time, error) {

	m.mu.RLock()
	sess, ok := m.sessions[token]
	m.mu.RUnlock()

	if !ok {
		return 0, time.Time{}, errors.New("session not found")
	}

	return sess.UserID, sess.ExpiresAt, nil
}

func (m *MemoryStore) DeleteSession(
	ctx context.Context, 
	token string) error {
		
	m.mu.Lock()
	delete(m.sessions, token)
	m.mu.Unlock()
	return nil
}