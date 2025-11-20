package user

import (
	"context"
	"time"

	domain "github.com/Ftthreign/aegisbox/backend/internal/domain/user"
)

type service struct {
	repo    Repository
	hasher  PasswordHasher
	keygen  KeyGenerator
	policy  Policy
}

func New(repo Repository, hasher PasswordHasher, keygen KeyGenerator, policy Policy) Service {
	return &service{repo: repo, hasher: hasher, keygen: keygen, policy: policy}
}

func (s *service) Register(ctx context.Context, email, password string) (*UserDTO, error) {
	existing, _ := s.repo.FindByEmail(ctx, email)
	if existing != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	if len(password) < s.policy.MinPasswordLength() {
		return nil, domain.ErrInvalidCredentials
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	salt, err := s.keygen.GenerateSalt(32)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        email,
		PasswordHash: hash,
		KeySalt:      salt,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &UserDTO{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (s *service) Authenticate(ctx context.Context, email, password string) (*UserDTO, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, domain.ErrInvalidCredentials
	}

	now := time.Now().UTC()

	if user.LockedUntil != nil && now.Before(*user.LockedUntil) {
		return nil, domain.ErrAccountLocked
	}

	if err := s.hasher.Verify(user.PasswordHash, password); err != nil {
		_ = s.IncrementFailed(ctx, user.ID)
		return nil, domain.ErrInvalidCredentials
	}

	_ = s.ResetFailed(ctx, user.ID)

	user.LastLogin = &now
	_ = s.repo.Update(ctx, user)

	return &UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		LastLogin: user.LastLogin,
	}, nil
}

func (s *service) UpdateLastLogin(ctx context.Context, id uint, t time.Time) error {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	u.LastLogin = &t
	return s.repo.Update(ctx, u)
}

func (s *service) IncrementFailed(ctx context.Context, id uint) error {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	u.FailedLoginAttempts++

	if u.FailedLoginAttempts >= s.policy.MaxFailedLogin() {
		lock := time.Now().Add(s.policy.LockDuration())
		u.LockedUntil = &lock
	}

	return s.repo.Update(ctx, u)
}

func (s *service) ResetFailed(ctx context.Context, id uint) error {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	u.FailedLoginAttempts = 0
	u.LockedUntil = nil

	return s.repo.Update(ctx, u)
}
