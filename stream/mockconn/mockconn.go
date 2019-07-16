package conn

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"time"
)

type WebsocketConn struct {
	conn uint
}

func (webConn *WebsocketConn) SetReadDeadline(t time.Time) error {
	//err := webConn.conn.SetReadDeadline(t)
	return nil
}

func (webConn *WebsocketConn) SetWriteDeadline(t time.Time) error {
	//err := webConn.conn.SetWriteDeadline(t)
	return nil
}

func (webConn *WebsocketConn) SetPongHandler(h func(appData string) error) {
	//webConn.conn.SetPongHandler(h)
}

func (webConn *WebsocketConn) MakeConn(ctx echo.Context, upgrader *websocket.Upgrader) error {
	//c, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	//webConn.conn = c
	return nil
}

func (webConn *WebsocketConn) Close() {
	//webConn.conn.Close()
}

func (webConn *WebsocketConn) WriteMessage(n int, p []byte) error {
	//err := webConn.conn.WriteMessage(n, p)
	return nil
}

func (webConn *WebsocketConn) ReadMessage() ([]byte, error) {
	//_, bytes, err := webConn.conn.ReadMessage()
	return "Test", nil
}
