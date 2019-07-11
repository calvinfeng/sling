// Helper functions to access/update db.
package model

import (
	"errors"
	"github.com/jinzhu/gorm" 
	"fmt"
	"strconv"
	"time"
	"github.com/golang-migrate/migrate"
	"github.com/calvinfeng/sling/model"
	
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	_ "github.com/lib/pq" // Driver
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver

)

//GetUsersInARoom : get all users for a particular room.
func GetUsersInARoom (db *gorm.DB, roomID string) []*model.User, error {
	sqlStatement := `
	SELECT users.id, users.name
	FROM userroom, users
	WHERE room_id = ?;`
	
	rows, err := db.Raw(sqlStatement, roomID).Rows()
	if err != nil {
		return [], err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.UserName)
		if err != nil {
		// handle this error
			return [],err
		}
		users = append(users, user)
	}
	return users, nil
}


//GetRooms: get all rooms, populating whether given user has joined and notification status.
func GetRooms (db *gorm.DB, userID string) []*model.Room, error {
	sqlStatement := `
	SELECT rooms.id, rooms.name, usersrooms.user_id IS NULL as inroom, 
	COALESCE(usersrooms.unread, false), rooms.room_type
	FROM rooms LEFT JOIN usersrooms
	ON rooms.id = usersrooms.room_id
	WHERE usersrooms.user_id=? OR usersrooms.user_id IS NULL;`
	
	rows, err := db.Row(sqlStatement, userID).Rows()
	if err != nil {
		return [], err
	}
	defer rows.Close()
	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		err = rows.Scan(&room.ID, &room.RoomName, &room.HasUnread)
		if err != nil {
		// handle this error
			return [],err
		}
		rooms = append(rooms, &room)
	}
	return rooms,nil
}

// GetAllMessagesFromRoom get all messages for a particular room.
func GetAllMessagesFromRoom (db *gorm.DB, roomID string) []model.Message, error {
	rows, err := db.Table("messages")
		.Select("messages.id, messages.time, users.name, messages.body, messages.sender_id, messages.room_id")
		.Joins("join users on messages.sender_id = users.id")
		.Where("messages.room_id = ?", roomID).Rows()
	if err != nil {
		return [], err
	}
	defer rows.Close()
	var messages []model.Message
	for rows.Next() {
		var message model.Message
		err = rows.Scan(&message)
		if err != nil {
			return [],err
		}
		messages = append(messages, message)
	}
	return messages,nil
}

// UpdateNotificationStatus update where a user has notification from given room
func UpdateNotificationStatus (db *gorm.DB, roomID string, userID string, hasUnread bool) error {
	sqlStatement := `
	UPDATE usersrooms
	SET unread = ?
	WHERE room_id = ? AND user_id = ?;`
	return db.Exec(sqlStatement, hasUnread, roomID, userID).Error
}

// GetUserNameByID: get a user's name by id
func GetUserNameByID (db *gorm.DB, user_id string) string {
	var userName string
	db.Table("users").Select("name").Where("id = ?", user_id).Scan(&userName)
	return userName
}

// InsertUser: insert a new user
func InsertUser (db *gorm.DB, u model.User) error {
	sqlStatement := `
	INSERT INTO users (name, email, jwt_token, password_digest)
	VALUES(?,?,?,?)`
	return db.Exec(sqlStatement, u.Name, u.Email, u.JWTToken, u.PasswordDigest).Error
}

// InsertRoom creates a new room and returns its ID as string
func InsertRoom (db *gorm.DB, room_name string, room_type int) string, error {
	newRoom := &model.Room{
		Name: room_name,
		RoomType: room_type,
	}
	err := db.Create(newRoom)
	if err != nil {
		return nil, err
	}
	return strconv.Itoa(newRoom.ID), nil
}

// InsertDMRoom creates a new direct message room and returns its ID and name as string
func InsertDMRoom (db *gorm.DB, user_id string, tar_user_id string) string, string, error {
	var userName, tarUserName string
	userName := GetUserNameByID(user_id)
	tarUserName := GetUserNameByID(tar_user_id)
	if userName == "" || tarUserName == "" {
		return nil, errors.New("User not found or empty name")
	}
	newRoomName := userName + "~" + tarUserName
	newRoomID, err : = InsertRoom(newRoomName, 1)
	if err != nil {
		return nil, nil, err
	}
	return newRoomId, newRoomName, nil
}

// InsertUserroom insert a new userroom
func InsertUserroom (db *gorm.DB, user_id string, room_id string, unread bool) error {
	sqlStatement := `
	INSERT INTO usersrooms (user_id, room_id, unread)
	VALUES(?,?,?)`
	return db.Exec(sqlStatement, user_id, room_id, unread).Error
} 

// InsertMessage insert a new message
func InsertMessage (db *gorm.DB, time time.Time, body string, sender_id string, room_id string) error {
	sqlStatement := `
	INSERT INTO messages (time, body, sender_id, room_id)
	VALUES(?,'?',?,?)`
	return db.Exec(sqlStatement, time).Error
}