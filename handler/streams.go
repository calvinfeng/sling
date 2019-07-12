/*==============================================================================
stream.go - Websocket Stream Interfaces
Summary: handles websocket connection request for actions and messages. Both
handlers use a channel list in broker (MessageBroker) to sync and create client.
==============================================================================*/

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/calvinfeng/sling/util"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// GetActionStreamHandler : handles the websocket connection request for
// actions. sends a reference to its' connection to the message websocket
// handler using a channel stored by the message broker, mapped by userID
func GetActionStreamHandler(upgrader *websocket.Upgrader) echo.HandlerFunc {
	if broker == nil {
		util.LogErr("please run you broker with RunBroker", nil)
		return nil
	}

	return func(ctx echo.Context) error {
		util.LogInfo("initiate action connection")
		actionConn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			util.LogErr("failure to upgrade action request - closing connection", err)
			actionConn.Close()
			return nil
		}

		_, bytes, err := actionConn.ReadMessage()
		util.LogInfo("read action connection")

		c := &TokenCredential{}
		errM := json.Unmarshal(bytes, c) // converts json to payload
		if errM != nil {
			util.LogErr("Error in reading token - closing connection", errM)
			actionConn.Close()
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByToken(broker.db, c.JWTToken)
		if err != nil {
			actionConn.Close()
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}

		// make host channel for other socket to pass connection to
		broker.mux.Lock()
		ch := make(chan *websocket.Conn)
		broker.websocketsByUserID[user.ID] = ch
		broker.mux.Unlock()
		// send actionConn along this channel
		broker.websocketsByUserID[user.ID] <- actionConn

		return nil
	}

}

// GetMessageStreamHandler : handles the websocket connection request for
// messages. waits for a reference to the action websocket's connection
// to be passed along a channel stored by the message broker, mapped by userID
func GetMessageStreamHandler(upgrader *websocket.Upgrader) echo.HandlerFunc { // handle message streams?
	if broker == nil {
		util.LogErr("please run you broker with RunBroker", nil)
		return nil
	}

	return func(ctx echo.Context) error {
		util.LogInfo("initiate message connection")

		messageConn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			util.LogErr("failure to upgrade action request - closing connection", err)
			messageConn.Close()
			return nil
		}

		_, bytes, err := messageConn.ReadMessage()
		util.LogInfo("read message connection")

		c := &TokenCredential{}

		errM := json.Unmarshal(bytes, c) // converts json to payload
		util.LogInfo("umarsh message connection")

		if errM != nil {
			util.LogErr("Error in reading token - closing connection", errM)
			messageConn.Close()
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByToken(broker.db, c.JWTToken)
		if err != nil {
			messageConn.Close()
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}

		var actionConn *websocket.Conn = getActionConn(user.ID)

		connectClient(messageConn, actionConn, user.ID)

		return nil
	}

}

// connectClient : creates a webSocketClient and adds both connection to it.
// Begins listening routines on both connectins, and waits until both are done.
func connectClient(messageConn *websocket.Conn, actionConn *websocket.Conn, userID uint) {
	defer messageConn.Close() // defer to execute after return
	defer actionConn.Close()

	cli := newWebSocketClient(messageConn, actionConn, userID)

	broker.addClient <- cli //this will activate read/write loops

	defer func() {
		broker.removeClient <- cli
	}()

	util.LogInfo(fmt.Sprintf("client %d has joined the chatroom", cli.UserID()))

	var wg sync.WaitGroup
	wg.Add(2)
	go cli.MessageListen(&wg)
	go cli.ActionListen(&wg)
	wg.Wait()

	util.LogInfo(fmt.Sprintf("client %d has left the chatroom", cli.UserID()))
}

// getActionConn : waits for action stream connection to be passed along channel
// mapped by userID, and then returns a reference to that connection
func getActionConn(userID uint) *websocket.Conn {
	// TODO: set a timeout
	for {
		// wait for this users channel to be created
		broker.mux.Lock()
		ch, ok := broker.websocketsByUserID[userID]
		broker.mux.Unlock()

		if ok {
			// wait for the connection information to be sent
			for {
				select {
				case actionConn := <-ch:
					return actionConn
				default:
				}
			}
		}
	}
}
