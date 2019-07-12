/*==============================================================================
payloads.go - Defines Payload Structures
Summary: Defines payloads for channel communication ingoing and outgoing between
the MessageBroker and WebsocketClient s
==============================================================================*/

package handler

import (
	"github.com/calvinfeng/sling/model"
	"time"
)

/************* Client to MessageBroker payload Types ******************/

// MessagePayload holds the message content to be communicated from the
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
