package user

import "time"

type User struct {
	ID                  uint
	Email               string
	PasswordHash        string
	KeySalt             []byte
	MFASecret           *string
	MFAEnabled          bool
	MFABackupCodes      []string
	FailedLoginAttempts int
	LockedUntil         *time.Time
	LastLogin           *time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
