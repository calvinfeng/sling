package model

// Room is a model for rooms.
type Room struct {
	ID       uint   `gorm:"column:id"        json:"id"`
	RoomName string `gorm:"column:room_name"   json:"room_name"`
	Type 	 int    `gorm:"column:room_type"   json:"room_type"`
	HasUnread bool  `gorm:"column:has_unread"   json:"has_unread"`
}

// TableName tells GORM where to find this record.
func (Room) TableName() string {
	return "rooms"
}

// Validate performs validation on message model.
func (r *Room) Validate() error {
	if len(r.RoomName) == 0 {
		return &ValidationError{Field: "name", Message: "cannot be empty"}
	}
	return nil
}
