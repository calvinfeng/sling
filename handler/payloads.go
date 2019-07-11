/*==============================================================================
payloads.go - Defines Payload Structures

Summary: Defines payloads for channel communication ingoing and outgoing between
the MessageBroker and WebsocketClient s
==============================================================================*/

package handler

/************* Client to MessageBroker payload Types ******************/

// MessagePayload holds the message content to be communicated from the
// client frontend, to the server and message broker
type MessagePayload struct {
	userID string `json:"userID"`
	roomID string `json:"roomID"`
	time   string `json:"time"`
	body   string `json:"body"`
}

// ActionPayload holds the action content to be communicated from the
// client frontend, to the server and message broker
type ActionPayload struct {
	actionType  string `json:"actionType"`
	userID      string `json:"userID"`
	roomID      string `json:"roomID"`
	newRoomID   string `json:"newRoomID"`
	dmUserID    string `json:"dmUserID"`
	newRoomName string `json:"newRoomName"`
}

/***************** MessageBroker to Client payload Types *****************/

// MessageResponsePayload holds the message content to be communicated from the
// message broker to users logged on
type MessageResponsePayload struct {
	messageType string `json:"messageType"`
	userName    string `json:"userName"`
	roomID      string `json:"roomID"`
	time        string `json:"time"`
	body        string `json:"body"`
}

// ActionResponsePayload holds the message content to be communicated from the
// message broker to users logged on
type ActionResponsePayload struct {
	actionType     string           `json:"actionType"`
	userID         string           `json:"userID"`
	roomID         string           `json:"roomID"`
	userName       string           `json:"userName"`
	roomName       string           `json:"roomName"`
	messageHistory []*model.Message `json:"messageHistory"`
}
