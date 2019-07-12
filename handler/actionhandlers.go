/*==============================================================================
actionhandlers.go - Additional Message Broker Functions: Action Responses

Summary: includes all handlers for actions change_room, send_message, create_dm,
 join_room, create_user, and create_room. Performs database changes & broadasts
==============================================================================*/
//NOTE: all database commands are not completed, and are marked with "DATABASE"

package handler

import (
	"fmt"

	"github.com/calvinfeng/sling/model"
	"github.com/calvinfeng/sling/util"
	"github.com/jinzhu/gorm"
)

func (mb *MessageBroker) handleChangeRoom(p ActionPayload) {
	// update group by room id CONSIDER: ordering of this update and the database
	cli := mb.clientByID[p.UserID]
	cli.SetRoomID(p.NewRoomID)
	if mb.groupByRoomID[p.NewRoomID] == nil {
		mb.groupByRoomID[p.NewRoomID] = make(map[uint]Client)
	}
	delete(mb.groupByRoomID[p.RoomID], p.UserID)
	mb.groupByRoomID[p.NewRoomID][p.UserID] = cli

	go func(db *gorm.DB, p ActionPayload, cli Client) {
		err := model.UpdateNotificationStatus(db, p.NewRoomID, p.UserID, false)
		if err != nil {
			util.LogErr("Error updating notification status", err)
		}

		messageHistory, err := model.GetAllMessagesFromRoom(db, p.NewRoomID)
		if err != nil {
			util.LogErr("failure to fetch message history", err)
			return
		}

		responsePayload := ActionResponsePayload{
			ActionType:     "message_history",
			MessageHistory: messageHistory,
		}

		cli.WriteActionQueue() <- responsePayload
	}(mb.db, p, cli)
}

func (mb *MessageBroker) handleCreateDm(p ActionPayload) {
	// create room first
	roomID, roomName, err := model.InsertDMRoom(mb.db, p.UserID, p.DMUserID)
	if err != nil {
		util.LogErr("could not insert new dm room", err)
		return
	}
	// update group by RoomID
	cli := mb.clientByID[p.UserID]
	if mb.groupByRoomID[roomID] == nil {
		mb.groupByRoomID[roomID] = make(map[uint]Client)
	}
	if mb.groupByRoomID[cli.RoomID()] != nil {
		delete(mb.groupByRoomID[cli.RoomID()], p.UserID)
	}
	cli.SetRoomID(roomID)
	mb.groupByRoomID[roomID][p.UserID] = cli

	reciever, recieverActive := mb.clientByID[p.DMUserID]
	responsePayload := ActionResponsePayload{
		ActionType: "create_dm",
		RoomID:     roomID,
		RoomName:   roomName,
		UserID:     p.UserID,
	}

	go func(responsePayload ActionResponsePayload, sender Client, reciever Client, recieverActive bool) {
		// send new dm response to self to inform frontend of room name/id
		sender.WriteActionQueue() <- responsePayload
		// send new dm notification to target user if logged on
		if recieverActive {
			reciever.WriteActionQueue() <- responsePayload
		}
	}(responsePayload, cli, reciever, recieverActive)
}

func (mb *MessageBroker) handleJoinRoom(p ActionPayload) {
	// update group by RoomID
	cli := mb.clientByID[p.UserID]
	if mb.groupByRoomID[p.NewRoomID] == nil {
		mb.groupByRoomID[p.NewRoomID] = make(map[uint]Client)
	}
	if mb.groupByRoomID[cli.RoomID()] != nil {
		delete(mb.groupByRoomID[cli.RoomID()], p.UserID)
	}
	cli.SetRoomID(p.NewRoomID)
	mb.groupByRoomID[p.NewRoomID][p.UserID] = cli

	go func(db *gorm.DB, p ActionPayload, cli Client) {
		model.InsertUserroom(db, p.UserID, p.NewRoomID, false)
		messageHistory, err := model.GetAllMessagesFromRoom(db, p.RoomID)
		if err != nil {
			util.LogErr("unable to fetch message history", err)
			return
		}

		responsePayload := ActionResponsePayload{
			ActionType:     "message_history",
			MessageHistory: messageHistory,
		}

		cli.WriteActionQueue() <- responsePayload
	}(mb.db, p, cli)
}

func (mb *MessageBroker) handleCreateUser(p ActionPayload) {
	// make a copy of shared map structure
	clientByIDCopy := make(map[uint]Client)
	for key, value := range mb.clientByID {
		clientByIDCopy[key] = value
	}
	// handle database updates and channel communication in another routine
	go func(db *gorm.DB, p ActionPayload, clientByIDCopy map[uint]Client) {
		userName := model.GetUserNameByID(db, p.UserID)

		responsePayload := ActionResponsePayload{
			ActionType: "new_user",
			UserID:     p.UserID,
			UserName:   userName,
		}
		// broadcast new user message to all users logged on
		for _, cli := range clientByIDCopy {
			cli.WriteActionQueue() <- responsePayload
		}
	}(mb.db, p, clientByIDCopy)
}

func (mb *MessageBroker) handleCreateRoom(p ActionPayload) {
	// make a copy of shared map structure
	clientByIDCopy := make(map[uint]Client)
	for key, value := range mb.clientByID {
		clientByIDCopy[key] = value
	}
	// handle database updates and channel communication in another routine
	func(db *gorm.DB, p ActionPayload, clientByIDCopy map[uint]Client) {
		roomID, err := model.InsertRoom(db, p.NewRoomName, 0)
		if err != nil {
			return // TODO: Better error handling
		}
		model.InsertUserroom(db, p.UserID, roomID, false)

		responsePayload := ActionResponsePayload{
			ActionType: "new_room",
			RoomID:     roomID,
			RoomName:   p.NewRoomName,
		}

		// broadcast new user message to all users logged on
		for _, cli := range clientByIDCopy {
			cli.WriteActionQueue() <- responsePayload
		}
	}(mb.db, p, clientByIDCopy)
}
