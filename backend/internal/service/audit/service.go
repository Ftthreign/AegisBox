package audit

import (
	"context"

	domain "github.com/Ftthreign/aegisbox/backend/internal/domain/audit"
)

type Repository interface {
	Create(ctx context.Context, a *domain.AuditLog) error
}

type Service interface {
	Log(ctx context.Context, userID *uint, action string, resourceType *string, resourceID *uint, success bool, meta map[string]any) error
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Log(ctx context.Context, userID *uint, action string, resourceType *string, resourceID *uint, success bool, meta map[string]any) error {
	a := &domain.AuditLog{
		UserID:       userID,
		Action:       action,
		ResourceType: resourceType,
		ResourceID:   resourceID,
		Meta:         meta,
		Success:      success,
	}
	return s.repo.Create(ctx, a)
}
