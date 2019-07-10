package model

// Usersrooms is a model for userrooms.
type Usersrooms struct {
	//ID        uint `gorm:"column:id"        json:"id"`
	UserID    uint `gorm:"column:user_id"        json:"user_id"`
	RoomID    uint `gorm:"column:room_id"        json:"room_id"`
	HasUnread bool `gorm:"column:unread"        json:"unread"`
}

// TableName tells GORM where to find this record.
func (Usersrooms) TableName() string {
	return "usersRooms"
}

// Validate performs validation on message model.
func (u *Usersrooms) Validate() error {
	if u.UserID == 0 {
		return &ValidationError{Field: "userID", Message: "required"}
	}

	if u.RoomID == 0 {
		return &ValidationError{Field: "roomID", Message: "required"}
	}

	return nil
}
