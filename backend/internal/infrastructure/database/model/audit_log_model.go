package model

import (
	"time"
)

type AuditLogModel struct {
	ID           uint           `gorm:"primaryKey"`
	UserID       *uint          `gorm:"index"`
	Action       string         `gorm:"size:64;not null;index"`
	ResourceType *string        `gorm:"size:64"`
	ResourceID   *uint
	Meta         map[string]any `gorm:"type:jsonb"`
	IPAddress    *string        `gorm:"size:45"`
	Success      bool           `gorm:"default:true"`
	CreatedAt    time.Time      `gorm:"index"`

	User *UserModel `gorm:"foreignKey:UserID"`
}

func (AuditLogModel) TableName() string {
	return "audit_logs"
}
