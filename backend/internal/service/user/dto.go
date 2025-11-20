package user

import "time"

type UserDTO struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	LastLogin *time.Time `json:"last_login,omitempty"`
}
