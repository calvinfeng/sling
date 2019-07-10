package model

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
	PasswordDigest []byte `gorm:"column:password_digest" json:"-"`
}

// TableName tells GORM where to find this record.
func (User) TableName() string {
	return "users"
}

// Validate validates a user model.
func (u *User) Validate() error {
	if len(u.Name) == 0 {
		return &ValidationError{Field: "Name", Message: "Username cannot be empty"}
	}

	if len(u.Email) == 0 {
		return &ValidationError{Field: "Email", Message: "Email cannot be empty"}
	}

	if len(u.Password) < 6 {
		return &ValidationError{Field: "Password", Message: "Password must be at least 6 characters"}
	}

	return nil
}
