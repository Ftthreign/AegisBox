package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID                  uint           `gorm:"primaryKey"`
	Email               string         `gorm:"uniqueIndex;not null"`
	PasswordHash        string         `gorm:"not null"`
	KeySalt             []byte         `gorm:"type:bytea;not null"`
	MFASecret           *string
	MFAEnabled          bool           `gorm:"default:false"`
	MFABackupCodes      []string       `gorm:"type:text[]"`
	FailedLoginAttempts int            `gorm:"default:0"`
	LockedUntil         *time.Time
	LastLogin           *time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`

	Credentials []CredentialModel `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	AuditLogs   []AuditLogModel   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (UserModel) TableName() string {
	return "users"
}
