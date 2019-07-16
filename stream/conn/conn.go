/*==============================================================================
conn.go - Implementation of a connection interface, using gorilla websockets
==============================================================================*/

package conn

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"time"
)

//WebsocketConn holds a pointer to a websocket
type WebsocketConn struct {
	conn *websocket.Conn
}

//SetReadDeadline sets a deadline for the next read to finish by t Time
func (webConn *WebsocketConn) SetReadDeadline(t time.Time) error {
	err := webConn.conn.SetReadDeadline(t)
	return err
}

//SetWriteDeadline sets a deadline for the next write to finish by t Time
func (webConn *WebsocketConn) SetWriteDeadline(t time.Time) error {
	err := webConn.conn.SetWriteDeadline(t)
	return err
}

//SetPongHandler sets the pong handler of the connection to function h
func (webConn *WebsocketConn) SetPongHandler(h func(appData string) error) {
	webConn.conn.SetPongHandler(h)
}

//MakeConn sets and creates all objects necessary for the internal webConn object
func (webConn *WebsocketConn) MakeConn(ctx echo.Context, upgrader *websocket.Upgrader) error {
	c, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	webConn.conn = c
	return err
}

//Close closes the websocket connection
func (webConn *WebsocketConn) Close() {
	webConn.conn.Close()
}

//Write Message tries to write a message to the websocket connection
func (webConn *WebsocketConn) WriteMessage(n int, p []byte) error {
	err := webConn.conn.WriteMessage(n, p)
	return err
}

//Write Message tries to read a message from the websocket connection
func (webConn *WebsocketConn) ReadMessage() ([]byte, error) {
	_, bytes, err := webConn.conn.ReadMessage()
	return bytes, err
}
