package handler

import (
	"encoding/json"
	"fmt"
	"github.com/calvinfeng/sling/util"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

// GetActionStreamHandler handles stream messages - to be called after initializing your broker
func GetActionStreamHandler(upgrader *websocket.Upgrader) echo.HandlerFunc { // handle message streams?
	if broker == nil {
		util.LogErr("please run you broker with RunBroker", nil)
		return nil
	}

	return func(ctx echo.Context) error {
		actionConn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			util.LogErr("failure to upgrade action request", err)
			return nil
		}

		_, bytes, err := actionConn.ReadMessage()

		c := &Credential{}
		errM := json.Unmarshal(bytes, c) // converts json to payload
		if errM != nil {
			util.LogErr("Error in reading token", errM)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByCredentials(broker.db, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}
		// make host channel for other socket to pass connection to
		broker.websocketsByUserID[user.ID] = make(chan *websocket.Conn)
		// send actionConn along this channel
		util.LogInfo("going to write to chan")

		broker.websocketsByUserID[user.ID] <- actionConn
		util.LogInfo("reached end")
		return nil
	}

}

// GetMessageStreamHandler handles stream messages - to be called after initializing your broker
func GetMessageStreamHandler(upgrader *websocket.Upgrader) echo.HandlerFunc { // handle message streams?
	if broker == nil {
		util.LogErr("please run you broker with RunBroker", nil)
		return nil
	}

	return func(ctx echo.Context) error {
		messageConn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			util.LogErr("failure to upgrade action request", err)
			return nil
		}

		_, bytes, err := messageConn.ReadMessage()

		c := &Credential{}
		errM := json.Unmarshal(bytes, c) // converts json to payload
		if errM != nil {
			util.LogErr("Error in reading token", errM)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByCredentials(broker.db, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}

		var actionConn *websocket.Conn = getActionConn(user.ID)

		connectClient(messageConn, actionConn, user.ID)
		return nil
	}

}

func connectClient(messageConn *websocket.Conn, actionConn *websocket.Conn, userID uint) {
	defer messageConn.Close() // defer to execute after return
	defer actionConn.Close()

	cli := newWebSocketClient(messageConn, actionConn, userID)

	broker.addClient <- cli //this will activate read/write loops

	defer func() { // defer to execute after return
		broker.removeClient <- cli
	}()

	util.LogInfo(fmt.Sprintf("client %s has joined the chatroom", cli.UserID()))

	var wg sync.WaitGroup
	wg.Add(2)
	go cli.MessageListen(wg)
	go cli.ActionListen(wg)
	wg.Wait()

	util.LogInfo(fmt.Sprintf("client %s has left the chatroom", cli.UserID()))
}

// getActionConn : waits for action stream connection to be passed along channel
// mapped by userID
func getActionConn(userID uint) *websocket.Conn {
	// TODO: set a timeout
	for {
		// wait for this users channel to be created
		ch, ok := broker.websocketsByUserID[userID]
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
