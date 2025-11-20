package credential

import (
	"context"

	domain "github.com/Ftthreign/aegisbox/backend/internal/domain/credential"
)

type Repository interface {
	Create(ctx context.Context, c *domain.Credential) error
	FindByUserID(ctx context.Context, uid uint) ([]domain.Credential, error)
	FindByID(ctx context.Context, id, uid uint) (*domain.Credential, error)
	Update(ctx context.Context, c *domain.Credential) error
	Delete(ctx context.Context, id, uid uint) error
}

type KeyDeriver interface {
	DeriveKey(password []byte, salt []byte) []byte
}

type Crypto interface {
	Encrypt(plaintext string, key []byte) (string, error)
	Decrypt(cipher string, key []byte) (string, error)
}

type UserReader interface {
	GetUserSalt(ctx context.Context, uid uint) ([]byte, error)
}

type Service interface {
	Create(ctx context.Context, dto *CredentialDTO) (*CredentialDTO, error)
	List(ctx context.Context, uid uint) ([]*CredentialDTO, error)
	Get(ctx context.Context, id, uid uint) (*CredentialDTO, error)
	Update(ctx context.Context, dto *CredentialDTO) error
	Delete(ctx context.Context, id, uid uint) error
}
