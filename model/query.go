// Helper functions to access/update db.
package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	_ "github.com/lib/pq"                                   // Driver
)

//GetUsersInARoom : get all users for a particular room.
func GetUsersInARoom(db *gorm.DB, roomID uint) ([]*User, error) {
	sqlStatement := `
	SELECT users.id, users.name
	FROM userroom, users
	WHERE room_id = ?;`

	rows, err := db.Raw(sqlStatement, roomID).Rows()
	if err != nil {
		return []*User{}, err
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		var user *User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			// handle this error
			return []*User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

//GetRooms: get all rooms, populating whether given user has joined and notification status.
func GetRooms(db *gorm.DB, userID uint) ([]*Room, error) {
	sqlStatement := `
	SELECT rooms.id, rooms.name, usersrooms.user_id IS NULL as inroom, 
	COALESCE(usersrooms.unread, false), rooms.room_type
	FROM rooms LEFT JOIN usersrooms
	ON rooms.id = usersrooms.room_id
	WHERE usersrooms.user_id=? OR usersrooms.user_id IS NULL;`

	rows, err := db.Exec(sqlStatement).Rows()
	if err != nil {
		return []*Room{}, err
	}
	defer rows.Close()
	var rooms []*Room
	for rows.Next() {
		var room *Room
		err = rows.Scan(&room.ID, &room.RoomName, &room)
		if err != nil {
			// handle this error
			return []*Room{}, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

// GetAllMessagesFromRoom get all messages for a particular room.
func GetAllMessagesFromRoom(db *gorm.DB, roomID uint) ([]*Message, error) {
	rows, err := db.Table("messages").
		Select("messages.id, messages.time, users.name, messages.body, messages.sender_id, messages.room_id").
		Joins("join users on messages.sender_id = users.id").
		Where("messages.room_id = ?", roomID).
		Rows()
	if err != nil {
		return []*Message{}, err
	}
	defer rows.Close()
	var messages []*Message
	for rows.Next() {
		var message *Message
		err = rows.Scan(&message)
		if err != nil {
			return []*Message{}, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

// UpdateNotificationStatus update where a user has notification from given room
func UpdateNotificationStatus(db *gorm.DB, roomID uint, userID uint, hasUnread bool) error {
	sqlStatement := `
	UPDATE usersrooms
	SET unread = ?
	WHERE room_id = ? AND user_id = ?;`
	return db.Exec(sqlStatement, hasUnread, roomID, userID).Error
}

// GetUserNameByID: get a user's name by id
func GetUserNameByID(db *gorm.DB, user_id uint) string {
	var userName string
	db.Table("users").Select("name").Where("id = ?", user_id).Scan(&userName)
	return userName
}

// InsertUser: insert a new user
func InsertUser(db *gorm.DB, u *User) error {
	sqlStatement := `
	INSERT INTO users (name, email, jwt_token, password_digest)
	VALUES(?,?,?,?)`
	return db.Exec(sqlStatement, u.Name, u.Email, u.JWTToken, u.PasswordDigest).Error
}

// InsertRoom creates a new room and returns its ID as string
func InsertRoom(db *gorm.DB, room_name string, room_type int) (uint, error) {
	newRoom := &Room{
		RoomName: room_name,
		Type:     room_type,
	}

	if dbc := db.Create(newRoom); dbc.Error != nil {
		return 0, dbc.Error
	}
	return newRoom.ID, nil
}

// InsertDMRoom creates a new direct message room and returns its ID and name as string
func InsertDMRoom(db *gorm.DB, user_id uint, tar_user_id uint) (uint, string, error) {
	userName, tarUserName := GetUserNameByID(db, user_id), GetUserNameByID(db, tar_user_id)
	if userName == "" || tarUserName == "" {
		return 0, "", errors.New("User not found or empty name")
	}
	newRoomName := userName + "~" + tarUserName
	newRoomID, err := InsertRoom(db, newRoomName, 1)
	if err != nil {
		return 0, "", err
	}
	return newRoomID, newRoomName, nil
}

// InsertUserroom insert a new userroom
func InsertUserroom(db *gorm.DB, user_id uint, room_id uint, unread bool) error {
	sqlStatement := `
	INSERT INTO usersrooms (user_id, room_id, unread)
	VALUES(?,?,?)`
	return db.Exec(sqlStatement, user_id, room_id, unread).Error
}

// InsertMessage insert a new message
func InsertMessage(db *gorm.DB, time time.Time, body string, sender_id uint, room_id uint) error {
	sqlStatement := `
	INSERT INTO messages (time, body, sender_id, room_id)
	VALUES(?,'?',?,?)`
	return db.Exec(sqlStatement, time).Error
}
