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
)

func (mb *MessageBroker) handleChangeRoom(p ActionPayload) {
	util.LogInfo("trying to call handleChangeRoom")

	err := model.UpdateNotificationStatus(mb.db, p.NewRoomID, p.UserID, false)
	if err != nil {
		util.LogErr("Error updating notification status", err)
	}

	// update groupByRoomID
	cli := mb.clientByID[p.UserID]
	cli.SetRoomID(p.NewRoomID)
	if mb.groupByRoomID[p.NewRoomID] == nil {
		mb.groupByRoomID[p.NewRoomID] = make(map[uint]Client)
	}
	delete(mb.groupByRoomID[p.RoomID], p.UserID)
	mb.groupByRoomID[p.NewRoomID][p.UserID] = cli

	messageHistory, err := model.GetAllMessagesFromRoom(mb.db, p.NewRoomID)
	if err != nil {
		return // TODO: Better error handling
	}

	responsePayload := ActionResponsePayload{
		ActionType:     "message_history",
		MessageHistory: messageHistory,
	}

	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateDm(p ActionPayload) {
	roomID, roomName, err := model.InsertDMRoom(mb.db, p.UserID, p.DMUserID)
	if err != nil {
		return // TODO: Better error handling
	}

	responsePayload := ActionResponsePayload{
		ActionType: "create_dm",
		RoomID:     roomID,
		RoomName:   roomName,
		UserID:     p.UserID,
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

	// send new dm response to self to inform frontend to change room
	cli.WriteActionQueue() <- responsePayload

	// send new dm notification to target user if logged on
	if cli, ok := mb.clientByID[p.DMUserID]; ok {
		cli.WriteActionQueue() <- responsePayload
	}
}

func (mb *MessageBroker) handleJoinRoom(p ActionPayload) {
	model.InsertUserroom(mb.db, p.UserID, p.NewRoomID, false)
	messageHistory, err := model.GetAllMessagesFromRoom(mb.db, p.NewRoomID)
	if err != nil {
		return // TODO: Better error handling
	}

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
	userName := model.GetUserNameByID(mb.db, p.UserID)

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
	roomID, err := model.InsertRoom(mb.db, p.NewRoomName, 0)
	if err != nil {
		fmt.Println(err)
		return // TODO: Better error handling
	}
	model.InsertUserroom(mb.db, p.UserID, roomID, false)

	responsePayload := ActionResponsePayload{
		ActionType: "new_room",
		UserID:     p.UserID,
		RoomID:     roomID,
		RoomName:   p.NewRoomName,
	}

	// broadcast new user message to all users logged on
	for _, cli := range mb.clientByID {
		cli.WriteActionQueue() <- responsePayload
	}
}
