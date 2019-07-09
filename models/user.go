package models

import (
	"errors"
	"time"
)

// User represents a user.
type User struct {
	// Both
	ID       uint   `gorm:"column:id"          json:"id"`
	Name     string `gorm:"column:name"        json:"name" `
	Email    string `gorm:"column:email"       json:"email"`
	JWTToken string `gorm:"column:jwt_token"   json:"jwt_token,omitempty"`

	// JSON only
	Password string `sql:"-" json:"password,omitempty"`

	// Database only
	CreatedAt      time.Time `gorm:"column:created_at"      json:"-"`
	UpdatedAt      time.Time `gorm:"column:updated_at"      json:"-"`
	PasswordDigest []byte    `gorm:"column:password_digest" json:"-"`
}

// Validate validates a user model.
func (u *User) Validate() error {
	if len(u.Name) == 0 {
		return errors.New("Username cannot be empty")
	}

	if len(u.Email) == 0 {
		return errors.New("Email cannot be empty")
	}

	if len(u.Password) < 6 {
		return errors.New("Password must be at least 6 characters")
	}

	return nil
}
