package model

import (
	"time"

	"gorm.io/gorm"
)

type CredentialModel struct {
	ID                uint           `gorm:"primaryKey"`
	UserID            uint           `gorm:"not null;index"`
	ServiceName       string         `gorm:"not null;index"`
	Username          string         `gorm:"not null"`
	EncryptedPassword string         `gorm:"type:text;not null"`
	Label             *string
	Notes             *string         `gorm:"type:text"`
	IsFavorite        bool            `gorm:"default:false"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt  `gorm:"index"`

	User UserModel `gorm:"foreignKey:UserID"`
}

func (CredentialModel) TableName() string {
	return "credentials"
}
