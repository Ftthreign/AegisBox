package credential

import (
	"context"
	"time"

	domain "github.com/Ftthreign/aegisbox/backend/internal/domain/credential"
)

type service struct {
	repo    Repository
	crypto  Crypto
	derive  KeyDeriver
	users   UserReader
}

func New(repo Repository, crypto Crypto, derive KeyDeriver, users UserReader) Service {
	return &service{repo: repo, crypto: crypto, derive: derive, users: users}
}

func (s *service) Create(ctx context.Context, dto *CredentialDTO) (*CredentialDTO, error) {
	salt, err := s.users.GetUserSalt(ctx, dto.UserID)
	if err != nil {
		return nil, err
	}

	key := s.derive.DeriveKey([]byte(dto.Password), salt)
	encrypted, err := s.crypto.Encrypt(dto.Password, key)
	if err != nil {
		return nil, err
	}

	model := &domain.Credential{
		UserID:            dto.UserID,
		ServiceName:       dto.ServiceName,
		Username:          dto.Username,
		EncryptedPassword: encrypted,
		Label:             dto.Label,
		Notes:             dto.Notes,
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Now().UTC(),
	}

	if err := s.repo.Create(ctx, model); err != nil {
		return nil, err
	}

	dto.ID = model.ID
	dto.Password = ""
	return dto, nil
}

func (s *service) List(ctx context.Context, uid uint) ([]*CredentialDTO, error) {
	creds, err := s.repo.FindByUserID(ctx, uid)
	if err != nil {
		return nil, err
	}

	out := make([]*CredentialDTO, 0, len(creds))

	for _, c := range creds {
		out = append(out, &CredentialDTO{
			ID:          c.ID,
			UserID:      uid,
			ServiceName: c.ServiceName,
			Username:    c.Username,
			Label:       c.Label,
			Notes:       c.Notes,
		})
	}
	return out, nil
}

func (s *service) Get(ctx context.Context, id, uid uint) (*CredentialDTO, error) {
	cred, err := s.repo.FindByID(ctx, id, uid)
	if err != nil {
		return nil, err
	}

	salt, err := s.users.GetUserSalt(ctx, uid)
	if err != nil {
		return nil, err
	}

	key := s.derive.DeriveKey(nil, salt) // decrypt-only mode

	pass, err := s.crypto.Decrypt(cred.EncryptedPassword, key)
	if err != nil {
		return nil, err
	}

	return &CredentialDTO{
		ID:          cred.ID,
		UserID:      uid,
		ServiceName: cred.ServiceName,
		Username:    cred.Username,
		Password:    pass,
		Label:       cred.Label,
		Notes:       cred.Notes,
	}, nil
}

func (s *service) Update(ctx context.Context, dto *CredentialDTO) error {
	cur, err := s.repo.FindByID(ctx, dto.ID, dto.UserID)
	if err != nil {
		return err
	}

	cur.ServiceName = dto.ServiceName
	cur.Username = dto.Username
	cur.Label = dto.Label
	cur.Notes = dto.Notes
	cur.UpdatedAt = time.Now().UTC()

	if dto.Password != "" {
		salt, _ := s.users.GetUserSalt(ctx, dto.UserID)
		key := s.derive.DeriveKey([]byte(dto.Password), salt)
		encrypted, err := s.crypto.Encrypt(dto.Password, key)
		if err != nil {
			return err
		}
		cur.EncryptedPassword = encrypted
	}

	return s.repo.Update(ctx, cur)
}

func (s *service) Delete(ctx context.Context, id, uid uint) error {
	return s.repo.Delete(ctx, id, uid)
}
