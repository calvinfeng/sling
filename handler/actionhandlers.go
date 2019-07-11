/*==============================================================================
actionhandlers.go - Additional Message Broker Functions: Action Responses

Summary: includes all handlers for actions change_room, send_message, create_dm,
 join_room, create_user, and create_room. Performs database changes & broadasts
==============================================================================*/
//NOTE: all database commands are not completed, and are marked with "DATABASE"

import "github.com/zpl0310/database/model"

package handler

func (mb *MessageBroker) handleChangeRoom(p ActionPayload) {
	model.UpdateNotificationStatus(mb.db, p.newRoomID, p.userID, false)

	// update groupByRoomID
	delete(mb.groupByRoomID[p.roomID],[p.userID])
	cli = mb.clientByID[p.userID]
	cli.RoomID = p.newRoomID
	mb.groupByRoomID[p.newRoomID],[p.userID]

	// DATABASE fetch list of messages in p.newRoomID
	// let messageHistory = list of messages type *model.Message (from dataModel)
	//messageHistory = []*model.Message
	messageHistory := model.GetAllMessagesFromRoom(mb.db, p.newRoomID)

	responsePayload = ActionResponsePayload{
		actionType: "message_history",
		messageHistory: messageHistory
	}

	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateDm(p ActionPayload) {
	roomID, roomName := model.InsertDMRoom(mb.db, p.userID, p.dmUserID)
	model.InsertUserroom(mb.db, p.userID, roomID, false)
	model.InsertUserroom(mb.db, p.dmUserID, roomID, true)

	responsePayload = ActionResponsePayload{
		actionType: "create_dm",
		roomId: roomID,
		roomName: roomName
	}

	// send new dm notification to users logged on
	if cli, ok = mb.clientByID[p.dmUserID]; ok {
		cli.WriteActionQueue() <- responsePayload
	}
}

func (mb *MessageBroker) handleJoinRoom(p ActionPayload) {
	model.InsertUserroom(p.userID, p.roomID, false)

	messageHistory := model.GetAllMessagesFromRoom(mb.db, p.newRoomID)

	responsePayload = ActionResponsePayload{
		actionType: "message_history",
		messageHistory: messageHistory
	}

	cli = mb.clientByID[p.userID]
	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateUser(p ActionPayload) {
	// database is already updated from a user user being created
	userName = model.GetUserNameByID(mb.db, p.userID)

	responsePayload = ActionResponsePayload{
		actionType: "new_user",
		userID: p.userID
		userName: userName
	}

	// broadcast new user message to all users logged on
	for _,cli:= range mb.clientByID {
		cli.WriteActionQueue() <- responsePayload
	}
}

func (mb *MessageBroker) handleCreateRoom(p ActionPayload) {
	roomID := model.InsertRoom(mb.db, p.newRoomName, 0)
	model.InsertUserroom(mb.db, p.userID, roomID, false)

	responsePayload = ActionResponsePayload{
		actionType: "new_room",
		roomID: roomID
		roomName: p.newRoomName
	}

	// broadcast new user message to all users logged on
	for _,cli:= range mb.clientByID {
		cli.WriteActionQueue() <- responsePayload
	}
}
