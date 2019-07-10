package handler

import (
	"sync"
)

// GetActionStreamHandler handles stream messages - to be called after initializing your broker
func GetActionStreamHandler(upgrader *websocket.Upgrader) echo.HandlerFunc { // handle message streams?
	if broker == nil {
		return nil, errors.New("please run you broker with RunBroker")
	}

	return func(ctx echo.Context) {
		c = Credential{}
		if err := ctx.Bind(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByCredentials(db, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}

		actionConn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			return
		}

		// make host channel for other socket to pass connection to
		broker.websocketsByUserID[user.ID] = make(chan *websocket.Conn)
		// send actionConn along this channel
		broker.websocketsByUserID[user.ID] <- actionConn

	}, nil

}

// GetMessageStreamHandler handles stream messages - to be called after initializing your broker
func GetMessageStreamHandler(upgrader *websocket.Upgrader) echo.HandlerFunc { // handle message streams?
	if broker == nil {
		return nil, errors.New("please run you broker with RunBroker")
	}

	return func(ctx echo.Context) {
		c = Credential{}
		if err := ctx.Bind(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByCredentials(db, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}

		messageConn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			return
		}

		var actionConn *websocket.Conn = getActionConn(user.ID)

		connectClient(messageConn, actionConn, user.ID)
	}, nil

}

func connectClient(messageConn *websocket.Conn, actionConn *websocket.Conn, userID uint) {
	defer messageConn.Close() // defer to execute after return
	defer actionConn.Close()

	cli := newWebSocketClient(messageConn, actionConn, userID)

	broker.addClient <- cli //this will activate read/write loops

	defer func() { // defer to execute after return
		broker.removeClient <- cli
	}()

	util.LogInfo(fmt.Sprintf("client %s has joined the chatroom", cli.ID()))

	var wg sync.WaitGroup
	wg.Add(2)
	go cli.MessageListen(wg)
	go cli.ActionListen(wg)
	wg.Wait()

	util.LogInfo(fmt.Sprintf("client %s has left the chatroom", cli.ID()))
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
