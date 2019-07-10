package model

// Userroom is a model for userrooms.
type Userroom struct {
	ID        uint `gorm:"column:id"        json:"id"`
	UserID    uint `gorm:"column:user_id"        json:"user_id"`
	RoomID    uint `gorm:"column:room_id"        json:"room_id"`
	HasUnread bool `gorm:"column:has_unread"        json:"has_unread"`
}

// TableName tells GORM where to find this record.
func (Userroom) TableName() string {
	return "userroom"
}

// Validate performs validation on message model.
func (u *Userroom) Validate() error {
	if u.UserID == 0 {
		return &ValidationError{Field: "userID", Message: "required"}
	}

	if u.RoomID == 0 {
		return &ValidationError{Field: "roomID", Message: "required"}
	}

	return nil
}
