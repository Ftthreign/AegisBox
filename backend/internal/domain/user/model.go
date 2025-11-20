package user

import "time"

type User struct {
	ID uint
	Email string
	MFAEnabled bool
	FailedLoginAttempts int 
	LockedUntil *time.Time
	LastLogin *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}