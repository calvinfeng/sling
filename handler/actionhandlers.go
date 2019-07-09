/*==============================================================================
actionhandlers.go - Additional Message Broker Functions: Action Responses

Summary: includes all handlers for actions change_room, send_message, create_dm,
 join_room, create_user, and create_room. Performs database changes & broadasts
==============================================================================*/
//NOTE: all database commands are not completed, and are marked with "DATABASE"


package handler


func (mb *MessageBroker) handleChangeRoom(p ActionPayload) {
	// DATABASE update usersrooms to have no unread notifications on p.roomId, p.userId

	// update groupByRoomID
	delete(mb.groupByRoomID[p.roomID],[p.userID])
	cli = mb.clientByID[p.userID]
	cli.RoomID = p.newRoomID
	mb.groupByRoomID[p.newRoomID],[p.userID]

	// DATABASE fetch list of messages in p.newRoomID
	// let messageHistory = list of messages type *model.Message (from dataModel)
	messageHistory = []*model.Message

	responsePayload = ActionResponsePayload{
		actionType: "message_history",
		messageHistory: messageHistory
	}

	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateDm(p ActionPayload) {
	// DATABASE update rooms to have new room of type dm with
	// users p.dmUserID and p.userID
	// DATABASE update usersrooms to mark new room as unread

	// return the new roomID and roomName
	roomID = "roomID"
	roomName = "roomName"

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
	// DATABASE update usersrooms to have room p.newRoomID and 
	// p.userID, read

	// DATABASE fetch list of messages in p.newRoomID
	// let messageHistory = list of messages type *model.Message (from dataModel)
	messageHistory = []*model.Message

	responsePayload = ActionResponsePayload{
		actionType: "message_history",
		messageHistory: messageHistory
	}

	cli = mb.clientByID[p.userID]
	cli.WriteActionQueue() <- responsePayload
}

func (mb *MessageBroker) handleCreateUser(p ActionPayload) {
	// database is already updated from a user user being created
	responsePayload = ActionResponsePayload{
		actionType: "message_history",
		messageHistory: messageHistory
	}
	
	// broadcast new user message to all users logged on
	for _,cli:= range mb.clientByID {
		cli.WriteActionQueue() <- responsePayload
	}

}

func (mb *MessageBroker) handleCreateRoom(p ActionPayload) {
	// TODO
}