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
			mb.handleSendMessage(b)
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
	// CONSIDER: ensure empty maps are deleted for memory saving?
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
	// make a copy of shared data
	clientsInRoom := make(map[uint]Client)
	clientsConnected := make(map[uint]Client)
	for key, value := range mb.groupByRoomID[p.RoomID] {
		clientsInRoom[key] = value
	}
	for key, value := range mb.clientByID {
		clientsConnected[key] = value
	}

	// go routine to handle lag of network calls and database calls
	go func(clientsConnected map[uint]Client, clientsInRoom map[uint]Client) {
		belongToRoom := make(map[uint]bool)

		model.InsertMessage(mb.db, p.Time, p.Body, p.UserID, p.RoomID)
		UsersInRoom, err := model.GetUsersInARoom(mb.db, p.RoomID)
		if err != nil {
			util.LogErr("users in room fetch err", err)
			return
		}
		for _, user := range UsersInRoom {
			if _, ok := clientsInRoom[user.ID]; !ok {
				model.UpdateNotificationStatus(mb.db, p.RoomID, user.ID, true)
			}
			belongToRoom[user.ID] = true
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
		for _, cli := range clientsInRoom {
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
			if cli, ok := clientsConnected[userID]; ok && active {
				select {
				case cli.WriteMessageQueue() <- notification:
				default:
				}
			}
		}
	}(clientsConnected, clientsInRoom)
}

// handleSendAction : checks action type, spawns appropriate goroutine handler
// NOTE: all action handler are in actionhandlers.go
func (mb *MessageBroker) handleSendAction(p ActionPayload) {
	switch p.ActionType {
	case "change_room":
		mb.handleChangeRoom(p)
	case "create_dm":
		mb.handleCreateDm(p)
	case "join_room":
		mb.handleJoinRoom(p)
	case "create_user":
		mb.handleCreateUser(p)
	case "create_room":
		mb.handleCreateRoom(p)
	}
}
