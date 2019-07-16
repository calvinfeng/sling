package stream

import (
	"context"
	"github.com/calvinfeng/sling/model"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"sync"
	"time"
)

//Client communicates with the Broker and its Conn connections
type Client interface {
	RoomID() uint
	UserID() uint
	MessageListen(*sync.WaitGroup)
	ActionListen(*sync.WaitGroup)
	Activate(ctx context.Context)
	WriteMessageQueue() chan<- MessageResponsePayload
	WriteActionQueue() chan<- ActionResponsePayload
	SetSendAction(chan ActionPayload)
	SetSendMessage(chan MessagePayload)
	SetRoomID(uint)
}

//Broker holds and manages all active client information
type Broker interface {
	SetSyncChannel(uint, chan Conn)
	GetSyncChannel(uint) (chan Conn, bool)
	HandleCreateUser(ActionPayload)
	LockMux()
	UnlockMux()
	GetDatabase() *gorm.DB
	AddClientQueue() chan Client
	RemoveClientQueue() chan Client
	CheckDuplicate(uint) bool
	DeleteSyncChannel(uint)
}

//Conn wraps any connection capabilities for different connections
type Conn interface {
	ReadMessage() ([]byte, error)
	WriteMessage(int, []byte) error
	Close()
	MakeConn(echo.Context, *websocket.Upgrader) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	SetPongHandler(func(string) error)
}

/************* Client to MessageBroker payload Types ******************/

//  stream.MessagePayload holds the message content to be communicated from the
// client frontend, to the server and message broker
type MessagePayload struct {
	UserID uint      `json:"userID"`
	RoomID uint      `json:"roomID"`
	Time   time.Time `json:"time"`
	Body   string    `json:"body"`
}

// ActionPayload holds the action content to be communicated from the
// client frontend, to the server and message broker
type ActionPayload struct {
	ActionType  string `json:"actionType"`
	UserID      uint   `json:"userID"`
	RoomID      uint   `json:"roomID"`
	NewRoomID   uint   `json:"newRoomID"`
	DMUserID    uint   `json:"dmUserID"`
	NewRoomName string `json:"newRoomName"`
}

/***************** MessageBroker to Client payload Types *****************/

// MessageResponsePayload holds the message content to be communicated from the
// message broker to users logged on
type MessageResponsePayload struct {
	MessageType string    `json:"messageType"`
	UserName    string    `json:"userName"`
	UserID      uint      `json:"userID"`
	RoomID      uint      `json:"roomID"`
	Time        time.Time `json:"time"`
	Body        string    `json:"body"`
}

// ActionResponsePayload holds the message content to be communicated from the
// message broker to users logged on
type ActionResponsePayload struct {
	ActionType     string                  `json:"actionType"`
	UserID         uint                    `json:"userID"`
	RoomID         uint                    `json:"roomID"`
	UserName       string                  `json:"userName"`
	RoomName       string                  `json:"roomName"`
	MessageHistory []*model.MessageHistory `json:"messageHistory"`
}
