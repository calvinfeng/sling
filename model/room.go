package model

// Room is a model for rooms.
type Room struct {
	ID       uint   `gorm:"column:id"        json:"id"`
	RoomName string `gorm:"column:name"   json:"name"`
	Type     int    `gorm:"column:room_type"        json:"room_type"`
}

// RoomDetail is a model for rooms with user-specific details,
// including whether it's unread and whether it's joined.
type RoomDetail struct {
	Room
	Joined bool `gorm:"column:inroom" json:"hasJoined"`
	Unread bool `gorm:"column:unread" json:"hasNotification"`
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
