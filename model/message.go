package model

import "time"

// Message is a model for messages.
type Message struct {
	ID         uint      `gorm:"column:id"          json:"id"`
	CreateTime time.Time `gorm:"column:time"        json:"time"`
	Body       string    `gorm:"column:body"        json:"body"`

	// Foreign keys
	SenderID uint `gorm:"column:sender_id"       json:"sender_id"`
	RoomID   uint `gorm:"column:room_id"       json:"room_id"`

	// Json only
	SenderName string `json:"sender_name"`
}

// TableName tells GORM where to find this record.
func (Message) TableName() string {
	return "messages"
}

// Validate performs validation on message model.
func (m *Message) Validate() error {
	if len(m.Body) == 0 {
		return &ValidationError{Field: "body", Message: "cannot be empty"}
	}

	if m.SenderID == 0 {
		return &ValidationError{Field: "userID", Message: "required"}
	}

	if m.RoomID == 0 {
		return &ValidationError{Field: "roomID", Message: "required"}
	}

	return nil
}
