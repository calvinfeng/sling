// Helper functions to access/update db.
package model

import (
	"errors"
	"fmt"
	"github.com/calvinfeng/sling/util"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	_ "github.com/lib/pq"                                   // Driver
)

//GetUsersInARoom : get all users for a particular room.
func GetUsersInARoom(db *gorm.DB, roomID uint) ([]*User, error) {
	rows, err := db.Select("users.id, users.name").
		Table("users").
		Joins("LEFT JOIN usersrooms ON users.id = usersrooms.user_id").
		Where("usersrooms.room_id = ?", roomID).Rows()

	if err != nil {
		return []*User{}, err
	}

	users := []*User{}

	defer rows.Close()
	for rows.Next() {
		user := &User{}
		err = db.ScanRows(rows, user)
		if err != nil {
			// handle this error
			return []*User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

//GetRooms: get all rooms, populating whether given user has joined and notification status.
func GetRooms(db *gorm.DB, userID uint) ([]*RoomDetail, error) {
	rooms := []*RoomDetail{}
	subquery := db.Select("usersrooms.*").Table("usersrooms").Where("user_id = ?", userID).SubQuery()
	rows, err := db.Select(`rooms.id, rooms.name, rooms.room_type,
		ur.user_id IS NOT NULL as inroom, COALESCE(ur.unread, false) as unread`).
		Table("rooms").
		Joins("LEFT JOIN ? as ur ON rooms.id = ur.room_id", subquery).
		Where("ur.user_id IS NOT NULL OR rooms.room_type = B'0'").
		Rows()

	if gorm.IsRecordNotFoundError(err) {
		return rooms, nil
	}

	if err != nil {
		return rooms, err
	}

	defer rows.Close()
	for rows.Next() {
		room := &RoomDetail{}
		err = db.ScanRows(rows, room)
		if err != nil {
			return []*RoomDetail{}, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

// GetAllMessagesFromRoom get all messages for a particular room.
func GetAllMessagesFromRoom(db *gorm.DB, roomID uint) ([]*MessageHistory, error) {
	rows, err := db.Table("messages").
		Select("messages.id, messages.time, messages.body, messages.sender_id, messages.room_id, users.name").
		Joins("join users on messages.sender_id = users.id").
		Where("messages.room_id = ?", roomID).
		Rows()

	if err != nil {
		return []*MessageHistory{}, err
	}

	messages := []*MessageHistory{}

	defer rows.Close()
	for rows.Next() {
		message := &MessageHistory{}
		err = db.ScanRows(rows, message)
		if err != nil {
			return []*MessageHistory{}, err
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
	userName := struct {
		Name string
	}{}
	db.Table("users").Select("name").Where("id = ?", user_id).Scan(&userName)
	return userName.Name
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

	// Add both members to the new room
	if err := InsertUserroom(db, user_id, newRoomID, false); err != nil {
		return 0, "", err
	}

	if err := InsertUserroom(db, tar_user_id, newRoomID, false); err != nil {
		return 0, "", err
	}

	return newRoomID, newRoomName, nil
}

// InsertUserroom insert a new userroom
func InsertUserroom(db *gorm.DB, user_id uint, room_id uint, unread bool) error {
	util.LogInfo(fmt.Sprintf("room id from query insert %d", room_id))
	sqlStatement := `
	INSERT INTO usersrooms (user_id, room_id, unread)
	VALUES(?,?,?)`
	return db.Exec(sqlStatement, user_id, room_id, unread).Error
}

// InsertMessage insert a new message
func InsertMessage(db *gorm.DB, time time.Time, body string, sender_id uint, room_id uint) error {
	sqlStatement := `
	INSERT INTO messages (time, body, sender_id, room_id)
	VALUES(?,?,?,?)`
	return db.Exec(sqlStatement, time, body, sender_id, room_id).Error
}
