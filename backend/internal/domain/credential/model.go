package credential

import "time"

type Credential struct {
	ID                uint
	UserID            uint
	ServiceName       string
	Username          string
	EncryptedPassword string
	Label             *string
	Notes             *string
	IsFavorite        bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
