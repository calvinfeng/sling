/*==============================================================================
stream.go - Websocket Stream Interfaces
Summary: handles websocket connection request for actions and messages. Both
handlers use a channel list in broker (MessageBroker) to sync and createstream.Client.
==============================================================================*/

package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/calvinfeng/sling/stream"
	"github.com/calvinfeng/sling/stream/client"
	"github.com/calvinfeng/sling/stream/conn"
	"github.com/calvinfeng/sling/util"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)

// GetActionStreamHandler : handles the websocket connection request for
// actions. sends a reference to its' connection to the message websocket
// handler using a channel stored by the message broker, mapped by userID
func GetActionStreamHandler(upgrader *websocket.Upgrader, broker stream.Broker) echo.HandlerFunc {
	if broker == nil {
		util.LogErr("please run you broker with RunBroker", nil)
		return nil
	}

	return func(ctx echo.Context) error {
		actionConn := &conn.WebsocketConn{}
		err := actionConn.MakeConn(ctx, upgrader)
		if err != nil {
			util.LogErr("failure to upgrade action request - closing connection", err)

			actionConn.Close()
			return nil
		}

		bytes, err := actionConn.ReadMessage()

		c := &TokenCredential{}
		errM := json.Unmarshal(bytes, c) // converts json to payload
		if errM != nil {
			util.LogErr("Error in reading token - closing connection", errM)
			actionConn.Close()
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByToken(broker.GetDatabase(), c.JWTToken)
		if err != nil {
			actionConn.Close()
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}
		if broker.CheckDuplicate(user.ID) {
			actionConn.Close()
			return echo.NewHTTPError(http.StatusUnauthorized, "this user is already logged in")
		}

		// make host channel for other socket to pass connection to
		broker.LockMux()
		ch := make(chan stream.Conn)
		broker.SetSyncChannel(user.ID, ch)
		broker.UnlockMux()

		// send actionConn along this channel
		errS := sendActionConn(actionConn, ch)
		if errS != nil {
			util.LogErr("websockets failed to sync", errS)
			actionConn.Close()
			return echo.NewHTTPError(http.StatusUnauthorized, "users could not connect")
		}

		return nil
	}

}

// GetMessageStreamHandler : handles the websocket connection request for
// messages. waits for a reference to the action websocket's connection
// to be passed along a channel stored by the message broker, mapped by userID
func GetMessageStreamHandler(upgrader *websocket.Upgrader, broker stream.Broker) echo.HandlerFunc { // handle message streams?
	if broker == nil {
		util.LogErr("please run you broker with RunBroker", nil)
		return nil
	}

	return func(ctx echo.Context) error {

		messageConn := &conn.WebsocketConn{}
		err := messageConn.MakeConn(ctx, upgrader)
		if err != nil {
			util.LogErr("failure to upgrade action request - closing connection", err)
			messageConn.Close()
			return nil
		}

		bytes, err := messageConn.ReadMessage()

		c := &TokenCredential{}

		errM := json.Unmarshal(bytes, c) // converts json to payload

		if errM != nil {
			util.LogErr("Error in reading token - closing connection", errM)
			messageConn.Close()
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByToken(broker.GetDatabase(), c.JWTToken)
		if err != nil {
			messageConn.Close()
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}
		if broker.CheckDuplicate(user.ID) {
			messageConn.Close()
			return echo.NewHTTPError(http.StatusUnauthorized, "this user is already logged in")
		}

		actionConn, err := getActionConn(broker, user.ID)
		if err != nil {
			util.LogErr("streams could not connect", err)
			messageConn.Close()
			return echo.NewHTTPError(http.StatusUnauthorized, "users could not connect")
		}
		connectClient(broker, messageConn, actionConn, user.ID)
		return nil
	}

}

// connectClient : creates a webSocketClient and adds both connection to it.
// Begins listening routines on both connectins, and waits until both are done.
func connectClient(broker stream.Broker, messageConn stream.Conn, actionConn stream.Conn, userID uint) {
	defer messageConn.Close() // defer to execute after return
	defer actionConn.Close()
	defer broker.DeleteSyncChannel(userID)

	cli := client.NewWebSocketClient(messageConn, actionConn, userID)

	broker.AddClientQueue() <- cli //this will activate read/write loops

	defer func() {
		broker.RemoveClientQueue() <- cli
	}()

	util.LogInfo(fmt.Sprintf("client %d has joined the chatroom", cli.UserID()))

	var wg sync.WaitGroup

	wg.Add(2)
	go cli.MessageListen(&wg)
	go cli.ActionListen(&wg)
	wg.Wait()

	util.LogInfo(fmt.Sprintf("client %d has left the chatroom", cli.UserID()))
}

// sendActionConn : sends connection along channel with timeout of 5 seconds
func sendActionConn(actionConn stream.Conn, ch chan stream.Conn) error {
	for {
		select {
		case ch <- actionConn:
			return nil
		case <-time.After(5 * time.Second):
			return errors.New("timeout on websocket sync")
		}
	}
}

// getActionConn : waits for action stream connection to be passed along channel
// mapped by userID, and then returns a reference to that connection
func getActionConn(broker stream.Broker, userID uint) (stream.Conn, error) {
	// set a timeout for 3 seconds
	timeoutErr := errors.New("timeout on websocket sync")

	for {
		// wait for this users channel to be created
		select {
		case <-time.After(2 * time.Second):
			return nil, timeoutErr
		default:
			broker.LockMux()
			ch, ok := broker.GetSyncChannel(userID)
			broker.UnlockMux()
			if ok {
				// wait for the connection information to be sent
				for {
					select {
					case actionConn := <-ch:
						return actionConn, nil
					case <-time.After(3 * time.Second):
						return nil, timeoutErr
					}
				}
			}
		}
	}
}
