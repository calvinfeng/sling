/*==============================================================================
broker.go - Core MessageBroker Functionality
Summary: Creates a MessageBroker, which handles messages along channels from
Websocket clients to perform broadcasts to other clients or change the database.
==============================================================================*/
//NOTE: all database commands are not completed, and are marked with "DATABASE"

package handler

import (
	"context"
	"sync"

	"github.com/calvinfeng/sling/model"
	"github.com/calvinfeng/sling/util"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	// "github.com/labstack/echo/v4"
	// _ "github.com/labstack/echo/v4/middleware"
)

// MessageBroker is a central hub that broadcasts all messages and actions,
// and sends commands to update the database
type MessageBroker struct {
	db                 *gorm.DB
	ctx                context.Context
	clientByID         map[uint]Client
	cancelByID         map[uint]context.CancelFunc
	sendMessage        chan MessagePayload
	addClient          chan Client
	removeClient       chan Client
	sendAction         chan ActionPayload
	groupByRoomID      map[uint]map[uint]Client
	websocketsByUserID map[uint]chan *websocket.Conn
	mux                *sync.Mutex
}

var broker *MessageBroker

// RunBroker : creates a broker and go routine to loop checking for messages
// from clients
func RunBroker(ctx context.Context, db *gorm.DB) {
	broker = &MessageBroker{
		db:                 db,
		ctx:                ctx,
		clientByID:         make(map[uint]Client),
		cancelByID:         make(map[uint]context.CancelFunc),
		sendMessage:        make(chan MessagePayload),
		addClient:          make(chan Client),
		removeClient:       make(chan Client),
		sendAction:         make(chan ActionPayload),
		groupByRoomID:      make(map[uint]map[uint]Client),
		websocketsByUserID: make(map[uint]chan *websocket.Conn),
		mux:                &sync.Mutex{},
	}

	go broker.loop()
}

// loop : spawns appropriate go routines for every new payload along a broker channel
func (mb *MessageBroker) loop() {
	for {
		select {
		case <-mb.ctx.Done():
			util.LogInfo("Stream: Done -- broker loop has terminated")
			return
		case c := <-mb.addClient:
			mb.handleAddClient(c)
		case c := <-mb.removeClient:
			mb.handleRemoveClient(c)
		case b := <-mb.sendMessage:
			go mb.handleSendMessage(b)
		case b := <-mb.sendAction:
			mb.handleSendAction(b)
		}
	}
}

// handleRemoveClient : removes client from broker structures, and cancels ctx
func (mb *MessageBroker) handleRemoveClient(c Client) {
	mb.cancelByID[c.UserID()]()
	delete(mb.cancelByID, c.UserID())
	delete(mb.clientByID, c.UserID())
	if c.RoomID() != 0 {
		delete(mb.groupByRoomID[c.RoomID()], c.UserID())
	}
	// TODO: ensure empty maps are deleted for memory saving?
}

// handleAddClient : creates client & child contexts, adds to broker structures
func (mb *MessageBroker) handleAddClient(c Client) {
	ctx, cancel := context.WithCancel(mb.ctx) // TODO : should I keep this as context.WithCancel?
	mb.clientByID[c.UserID()] = c
	mb.cancelByID[c.UserID()] = cancel

	c.SetSendMessage(mb.sendMessage)
	c.SetSendAction(mb.sendAction)
	c.Activate(ctx)
}

// handleSendMessage : updates database, sends messages and notifications to
// clients when the broker recieves a new message
func (mb *MessageBroker) handleSendMessage(p MessagePayload) {
	// DATABASE add message mb to database
	model.InsertMessage(mb.db, p.Time, p.Body, p.UserID, p.RoomID)

	// DATABASE for all users in room p.roomId,
	// whose Ids are not in mb.groupByRoomID[p.roomID], update to unread
	belongToRoom := make(map[uint]bool)
	UsersInRoom, err := model.GetUsersInARoom(mb.db, p.RoomID)
	curViewing := make(map[uint]bool)
	if err != nil {
		util.LogErr("users in room fetch err", err)
		return
	}

	for _, client := range mb.groupByRoomID[p.RoomID] {
		curViewing[client.UserID()] = true
	}

	for _, user := range UsersInRoom {
		if _, ok := curViewing[user.ID]; !ok {
			model.UpdateNotificationStatus(mb.db, p.RoomID, user.ID, true)
			belongToRoom[user.ID] = true
		}
	}

	name := model.GetUserNameByID(mb.db, p.UserID)

	// update p to be a notification type
	message := MessageResponsePayload{
		MessageType: "new_message",
		UserName:    name,
		UserID:      p.UserID,
		RoomID:      p.RoomID,
		Time:        p.Time,
		Body:        p.Body,
	}

	// send live messages to clients logged into this room
	for _, cli := range mb.groupByRoomID[p.RoomID] {
		select {
		case cli.WriteMessageQueue() <- message:
		default:
		}
		belongToRoom[cli.UserID()] = false
	}

	// update p to be a notification type
	notification := MessageResponsePayload{
		MessageType: "notification",
		RoomID:      p.RoomID,
	}

	// send live notifications to logged in clients who belong to this room
	for userID, active := range belongToRoom {
		if cli, ok := mb.clientByID[userID]; ok && active {
			select {
			case cli.WriteMessageQueue() <- notification:
			default:
			}
		}
	}
}

// handleSendAction : checks action type, spawns appropriate goroutine handler
// NOTE: all action handler are in actionhandlers.go
func (mb *MessageBroker) handleSendAction(p ActionPayload) {
	switch p.ActionType {
	case "change_room":
		go mb.handleChangeRoom(p)
	case "create_dm":
		go mb.handleCreateDm(p)
	case "join_room":
		go mb.handleJoinRoom(p)
	case "create_user":
		go mb.handleCreateUser(p)
	case "create_room":
		go mb.handleCreateRoom(p)
	}
}
