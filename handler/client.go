/*==============================================================================
Client.go - Websocket Client Interface

Summary: Stores information for each connected client in WebSocketClient,
Client interface allows communication with the MessageBroker, reads and writes
to all clients based off of channel communication.
==============================================================================*/

package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/calvinfeng/sling/util"
	"github.com/gorilla/websocket"

	// "github.com/labstack/echo/v4"
	"sync"
	"time"
)

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

type WebSocketClient struct {
	roomID       uint
	userID       uint
	connMessage  *websocket.Conn
	connAction   *websocket.Conn
	readMessage  chan json.RawMessage        // read next message
	writeMessage chan MessageResponsePayload // write to msg queue next
	readAction   chan json.RawMessage        // read next message
	writeAction  chan ActionResponsePayload  // write to msg queue next
	sendMessage  chan MessagePayload
	sendAction   chan ActionPayload
}

func newWebSocketClient(messageConn *websocket.Conn, actionConn *websocket.Conn, userID uint) Client {
	return &WebSocketClient{
		userID:       userID,
		roomID:       0,
		connMessage:  messageConn,
		connAction:   actionConn,
		readMessage:  make(chan json.RawMessage, 200),        // read next message
		writeMessage: make(chan MessageResponsePayload, 200), // write to msg queue next
		readAction:   make(chan json.RawMessage, 200),        // read next message
		writeAction:  make(chan ActionResponsePayload, 200),  // write to msg queue next
		sendMessage:  make(chan MessagePayload, 200),
		sendAction:   make(chan ActionPayload, 200),
	}
}

// UserID : returns userId, the user_id value in the database related to this client
func (c *WebSocketClient) UserID() uint {
	return c.userID
}

// RoomID : returns roomId, the room_id value in the database for this client
func (c *WebSocketClient) RoomID() uint {
	return c.roomID
}

func (c *WebSocketClient) SetRoomID(roomID uint) {
	c.roomID = roomID
}

// MessageListen : continously checks for new messages along the websocket connection,
// and forwards new messages along the proper channels
func (c *WebSocketClient) MessageListen(wg *sync.WaitGroup) {
	defer wg.Done()
	defer util.LogInfo(fmt.Sprintf("client %d disconnected from messages websocket", c.UserID()))

	for {
		_, bytes, err := c.connMessage.ReadMessage()

		if err != nil &&
			websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			return
		}

		if err != nil {
			util.LogErr("conn.ReadMessage", err)
			return
		}

		c.readMessage <- bytes // no errors; send the message to the read channel
	}
}

// ActionListen : continously checks for new messages along the websocket connection,
// and forwards new messages along the proper channels
func (c *WebSocketClient) ActionListen(wg *sync.WaitGroup) {
	defer wg.Done()
	defer util.LogInfo(fmt.Sprintf("client %d disconnected from actions websocket", c.UserID()))
	for {
		_, bytes, err := c.connAction.ReadMessage()

		if err != nil &&
			websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			return
		}

		if err != nil {
			util.LogErr("conn.ReadMessage", err)
			return
		}

		c.readAction <- bytes // no errors; send the message to the read channel
	}
}

// Activate : function creates a read and write go routine for this client.
func (c *WebSocketClient) Activate(ctx context.Context) {
	go c.readMessageLoop(ctx)
	go c.writeMessageLoop(ctx)
	go c.readActionLoop(ctx)
	go c.writeActionLoop(ctx)
}

// WriteMessageQueue : returns channel to write messages to
func (c *WebSocketClient) WriteMessageQueue() chan<- MessageResponsePayload {
	return c.writeMessage
}

// WriteActionQueue : returns channel to write actions to
func (c *WebSocketClient) WriteActionQueue() chan<- ActionResponsePayload {
	return c.writeAction
}

// SetSendMessage : sets channel for sending messages to message broker
func (c *WebSocketClient) SetSendMessage(ch chan MessagePayload) {
	c.sendMessage = ch
}

// SetSendAction : sets channel for sending actions to message broker
func (c *WebSocketClient) SetSendAction(ch chan ActionPayload) {
	c.sendAction = ch
}

// readMessageLoop : continuously reads from connMessage for new messages, and
// forwards them to the message broker
func (c *WebSocketClient) readMessageLoop(ctx context.Context) {
	c.connMessage.SetReadDeadline(time.Now().Add(2 * time.Second))
	c.connMessage.SetPongHandler(func(s string) error {
		c.connMessage.SetReadDeadline(time.Now().Add(2 * time.Second))
		return nil
	})

	for {
		select {
		case <-ctx.Done(): // read loop is closed
			util.LogInfo(fmt.Sprintf("client %d has terminated read message loop", c.userID))
			return

		case bytes := <-c.readMessage: // read bytes detected in channel
			c.connMessage.SetReadDeadline(time.Now().Add(2 * time.Second))
			p := MessagePayload{}
			util.LogInfo(string(bytes))

			err := json.Unmarshal(bytes, &p) // converts json to payload
			if err != nil {
				util.LogErr("readMessageLoop: json.Unmarshall", err)
				continue
			}
			c.sendMessage <- p
		}
	}
}

// readActionLoop : continuously reads from connAction for new actions, and
// forwards them to the message broker
func (c *WebSocketClient) readActionLoop(ctx context.Context) {
	c.connAction.SetReadDeadline(time.Now().Add(2 * time.Second))
	c.connAction.SetPongHandler(func(s string) error {
		c.connAction.SetReadDeadline(time.Now().Add(2 * time.Second))
		return nil
	})

	for {
		select {
		case <-ctx.Done(): // read loop is closed
			util.LogInfo(fmt.Sprintf("client %d has terminated read action loop", c.userID))
			return

		case bytes := <-c.readAction: // read bytes detected in channel
			c.connAction.SetReadDeadline(time.Now().Add(2 * time.Second))
			p := ActionPayload{}
			util.LogInfo(string(bytes))

			err := json.Unmarshal(bytes, &p) // converts json to payload
			if err != nil {
				util.LogErr("readActionLoop: json.Unmarshall", err)
				continue
			}
			c.sendAction <- p
		}
	}
}

// writeMessageLoop : writes messages to client if a message payload is passed
// along c.writeMessage
func (c *WebSocketClient) writeMessageLoop(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): // write loop is closed
			util.LogInfo(fmt.Sprintf("client %d has terminated write loop", c.userID))
			return
		case p := <-c.writeMessage: // message broker wants to write to client
			c.connMessage.SetReadDeadline(time.Now().Add(2 * time.Second))
			bytes, err := json.Marshal(p)
			if err != nil {
				util.LogErr("writeMessageLoop json.Marshall", err)
				continue
			}

			err = c.connMessage.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				return
			}

		case <-ticker.C: // send a ping to the connection if the ticker signals
			c.connMessage.SetWriteDeadline(time.Now().Add(5 * time.Second))
			err := c.connMessage.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		}
	}
}

// writeActionLoop : writes actions to client if a action payload is passed
// along c.writeAction
func (c *WebSocketClient) writeActionLoop(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): // write loop is closed
			util.LogInfo(fmt.Sprintf("client %d has terminated write loop", c.userID))
			return
		case p := <-c.writeAction: // message broker wants to write to client
			c.connAction.SetReadDeadline(time.Now().Add(2 * time.Second))
			bytes, err := json.Marshal(p)
			if err != nil {
				util.LogErr("writeActionLoop json.Marshall", err)
				continue
			}

			err = c.connAction.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				return
			}

		case <-ticker.C: // send a ping to the connection if the ticker signals
			c.connAction.SetWriteDeadline(time.Now().Add(5 * time.Second))
			err := c.connAction.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		}
	}
}
