package user

import (
	"context"
	"time"

	domain "github.com/Ftthreign/aegisbox/backend/internal/domain/user"
)

type Repository interface {
	Create(ctx context.Context, u *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	Update(ctx context.Context, u *domain.User) error
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(hashed string, password string) error
}

type KeyGenerator interface {
	GenerateSalt(length int) ([]byte, error)
}

type Policy interface {
	MinPasswordLength() int
	MaxFailedLogin() int
	LockDuration() time.Duration
}

type Service interface {
	Register(ctx context.Context, email, password string) (*UserDTO, error)
	Authenticate(ctx context.Context, email, password string) (*UserDTO, error)
	UpdateLastLogin(ctx context.Context, id uint, t time.Time) error
	IncrementFailed(ctx context.Context, id uint) error
	ResetFailed(ctx context.Context, id uint) error
}
