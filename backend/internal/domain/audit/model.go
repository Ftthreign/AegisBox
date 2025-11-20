package audit

import "time"

type AuditLog struct {
	ID           uint
	UserID       *uint
	Action       string
	ResourceType *string
	ResourceID   *uint
	Meta         map[string]any
	IPAddress    *string
	Success      bool
	CreatedAt    time.Time
}