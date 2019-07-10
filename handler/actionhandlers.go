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
	// DATABASE update usersrooms to have no unread notifications on p.roomId, p.userId

	// update groupByRoomID
	cli := mb.clientByID[p.userID]
	cli.SetRoomID(p.newRoomID)
	if mb.groupByRoomID[p.newRoomID] == nil {
		mb.groupByRoomID[p.newRoomID] = make(map[uint]Client)
	}
	delete(mb.groupByRoomID[p.roomID], p.userID)
	mb.groupByRoomID[p.newRoomID][p.userID] = cli

	// DATABASE fetch list of messages in p.newRoomID
	// let messageHistory = list of messages type *model.Message (from dataModel)
	var messageHistory []*model.Message

	responsePayload := ActionResponsePayload{
		actionType:     "message_history",
		messageHistory: messageHistory,
	}

	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateDm(p ActionPayload) {
	// DATABASE update rooms to have new room of type dm with
	// users p.dmUserID and p.userID
	// DATABASE update usersrooms to mark new room as unread

	// return the new roomID and roomName
	var roomID uint = 0
	roomName := "roomName"

	responsePayload := ActionResponsePayload{
		actionType: "create_dm",
		roomID:     roomID,
		roomName:   roomName,
	}

	// update group by roomID
	cli := mb.clientByID[p.userID]
	if mb.groupByRoomID[roomID] == nil {
		mb.groupByRoomID[roomID] = make(map[uint]Client)
	}
	if mb.groupByRoomID[cli.RoomID()] != nil {
		delete(mb.groupByRoomID[cli.RoomID()], p.userID)
	}
	cli.SetRoomID(roomID)
	mb.groupByRoomID[roomID][p.userID] = cli

	// send new dm notification to users logged on
	if cli, ok := mb.clientByID[p.dmUserID]; ok {
		cli.WriteActionQueue() <- responsePayload
	}
}

func (mb *MessageBroker) handleJoinRoom(p ActionPayload) {
	// DATABASE update usersrooms to have room p.newRoomID and
	// p.userID, read

	// DATABASE fetch list of messages in p.newRoomID
	// let messageHistory = list of messages type *model.Message (from dataModel)
	var messageHistory []*model.Message

	responsePayload := ActionResponsePayload{
		actionType:     "message_history",
		messageHistory: messageHistory,
	}

	// update group by roomID
	cli := mb.clientByID[p.userID]
	if mb.groupByRoomID[p.newRoomID] == nil {
		mb.groupByRoomID[p.newRoomID] = make(map[uint]Client)
	}
	if mb.groupByRoomID[cli.RoomID()] != nil {
		delete(mb.groupByRoomID[cli.RoomID()], p.userID)
	}
	cli.SetRoomID(p.newRoomID)
	mb.groupByRoomID[p.newRoomID][p.userID] = cli

	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateUser(p ActionPayload) {
	// database is already updated from a user user being created
	// DATABASE
	// let userName = fetch the user's name from the database
	userName := "userName"

	responsePayload := ActionResponsePayload{
		actionType: "new_user",
		userID:     p.userID,
		userName:   userName,
	}

	// broadcast new user message to all users logged on
	for _, cli := range mb.clientByID {
		cli.WriteActionQueue() <- responsePayload
	}
}

func (mb *MessageBroker) handleCreateRoom(p ActionPayload) {

	responsePayload := ActionResponsePayload{
		actionType: "new_user",
		roomID:     p.roomID,
		roomName:   p.newRoomName,
	}

	// broadcast new user message to all users logged on
	for _, cli := range mb.clientByID {
		cli.WriteActionQueue() <- responsePayload
	}
}
