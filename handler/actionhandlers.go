/*==============================================================================
actionhandlers.go - Additional Message Broker Functions: Action Responses

Summary: includes all handlers for actions change_room, send_message, create_dm,
 join_room, create_user, and create_room. Performs database changes & broadasts
==============================================================================*/
//NOTE: all database commands are not completed, and are marked with "DATABASE"

package handler

import (
	"github.com/calvinfeng/sling/model"
)

func (mb *MessageBroker) handleChangeRoom(p ActionPayload) {
	// DATABASE update usersrooms to have no unread notifications on p.roomId, p.UserID

	// update groupByRoomID
	cli := mb.clientByID[p.UserID]
	cli.SetRoomID(p.NewRoomID)
	if mb.groupByRoomID[p.NewRoomID] == nil {
		mb.groupByRoomID[p.NewRoomID] = make(map[uint]Client)
	}
	delete(mb.groupByRoomID[p.RoomID], p.UserID)
	mb.groupByRoomID[p.NewRoomID][p.UserID] = cli

	// DATABASE fetch list of messages in p.NewRoomID
	// let messageHistory = list of messages type *model.Message (from dataModel)
	var messageHistory []*model.Message

	responsePayload := ActionResponsePayload{
		ActionType:     "message_history",
		MessageHistory: messageHistory,
	}

	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateDm(p ActionPayload) {
	// DATABASE update rooms to have new room of type dm with
	// users p.dmUserID and p.UserID
	// DATABASE update usersrooms to mark new room as unread

	// return the new roomID and roomName
	var roomID uint
	roomName := "roomName"

	responsePayload := ActionResponsePayload{
		ActionType: "create_dm",
		RoomID:     roomID,
		RoomName:   roomName,
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

	// send new dm notification to users logged on
	if cli, ok := mb.clientByID[p.DMUserID]; ok {
		cli.WriteActionQueue() <- responsePayload
	}
}

func (mb *MessageBroker) handleJoinRoom(p ActionPayload) {
	// DATABASE update usersrooms to have room p.NewRoomID and
	// p.UserID, read

	// DATABASE fetch list of messages in p.NewRoomID
	// let MessageHistory = list of messages type *model.Message (from dataModel)
	var messageHistory []*model.Message

	responsePayload := ActionResponsePayload{
		ActionType:     "message_history",
		MessageHistory: messageHistory,
	}

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

	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateUser(p ActionPayload) {
	// database is already updated from a user user being created
	// DATABASE
	// let userName = fetch the user's name from the database
	userName := "userName"

	responsePayload := ActionResponsePayload{
		ActionType: "new_user",
		UserID:     p.UserID,
		UserName:   userName,
	}

	// broadcast new user message to all users logged on
	for _, cli := range mb.clientByID {
		cli.WriteActionQueue() <- responsePayload
	}
}

func (mb *MessageBroker) handleCreateRoom(p ActionPayload) {

	responsePayload := ActionResponsePayload{
		ActionType: "new_user",
		RoomID:     p.RoomID,
		RoomName:   p.NewRoomName,
	}

	// broadcast new user message to all users logged on
	for _, cli := range mb.clientByID {
		cli.WriteActionQueue() <- responsePayload
	}
}
