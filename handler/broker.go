/*==============================================================================
broker.go - Core MessageBroker Functionality

Summary: Creates a MessageBroker, which handles messages along channels from
Websocket clients to perform broadcasts to other clients or change the database.
==============================================================================*/
//NOTE: all database commands are not completed, and are marked with "DATABASE"

package handler

import (
	"context"
	_ "github.com/calvinfeng/sling/model"
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
	clientByID         map[string]Client
	cancelByID         map[string]context.CancelFunc
	sendMessage        chan MessagePayload
	addClient          chan Client
	removeClient       chan Client
	sendAction         chan ActionPayload
	groupByRoomID      map[string]map[string]Client
	websocketsByUserID map[string]chan *websocket.Conn
}

var broker *MessageBroker

// RunBroker : creates a broker and go routine to loop checking for messages
// from clients
func RunBroker(ctx context.Context, db *gorm.DB) {
	broker = &MessageBroker{
		db:                 db,
		ctx:                ctx,
		clientByID:         make(map[string]Client),
		cancelByID:         make(map[string]context.CancelFunc),
		sendMessage:        make(chan MessagePayload),
		addClient:          make(chan Client),
		removeClient:       make(chan Client),
		sendAction:         make(chan ActionPayload),
		groupByRoomID:      make(map[string]map[string]Client),
		websocketsByUserID: make(map[string]chan *websocket.Conn),
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
	if c.RoomID() != nil {
		delete(mb.groupByRoomID[c.RoomID()], c.UserID())
	}
	// TODO: ensure empty maps are deleted for memory saving?
}

// handleAddClient : creates client & child contexts, adds to broker structures
func (mb *MessageBroker) handleAddClient(c Client) {
	ctx, cancel := context.WithCancel(mb.messageCtx) // TODO : should I keep this as context.WithCancel?
	mb.clientByID[c.ID()] = c
	mb.cancelByID[c.ID()] = cancel

	// if mb.groupByRoomID[c.RoomID()] == nil {
	// 	mb.groupByRoomID[c.RoomID()] = make(map[string]Client)
	// }
	// mb.groupByRoomID[c.RoomID()][c.ID()] = c

	c.SetSendMessage(mb.sendMessage)
	c.SetSendAction(mb.sendAction)
	c.Activate(ctx)
}

// handleSendMessage : updates database, sends messages and notifications to
// clients when the broker recieves a new message
func (mb *MessageBroker) handleSendMessage(p MessagePayload) {
	// update database: TODO waiting for database implementations

	// DATABASE add message mb to database

	// DATABASE for all users in room p.roomId,
	// whose Ids are not in mb.groupByRoomID[p.roomID], update to unread

	// Let belongToRoom = map of user_ids to booleans that belong to p.roomID DATABASE
	belongToRoom = make(map[string]bool)

	// update p to be a notification type
	message = MessageResponsePayload{
		actionType: "new_message",
		userID:     p.userID,
		roomID:     p.roomID,
		time:       p.time,
		body:       p.body,
	}

	// send live messages to clients logged into this room
	for _, cli := range mb.groupByRoomID[p.roomID] {
		select {
		case cli.WriteMessageQueue() <- message:
		default:
		}
		belongToRoom[cli.userID()] = false
	}

	// update p to be a notification type
	notification = MessageResponsePayload{
		actionType: "notification",
		roomID:     p.roomID,
	}

	// send live notifications to logged in clients who belong to this room
	for userId, active := range belongToRoom {
		if cli, ok := clientByID[userId]; ok {
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
	select {
	case p.actionType == "change_room":
		go mb.handleChangeRoom(p)
	case p.actionType == "create_dm":
		go mb.handleCreateDm(p)
	case p.actionType == "join_room":
		go mb.handleJoinRoom(p)
	case p.actionType == "create_user":
		go mb.handleCreateUser(p)
	case p.actionType == "create_room":
		go mb.handleCreateRoom(p)
	}
}
