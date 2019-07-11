package model

// RoomDetail is a model for rooms with user-specific details,
// including whether it's unread and whether it's joined.
type RoomDetail struct {
	ID     uint   `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Type   int    `gorm:"column:room_type" json:"type"`
	Joined bool   `gorm:"column:inroom" json:"hasJoined"`
	Unread bool   `gorm:"column:unread" json:"hasNotification"`
}
