package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header, userId uint) (*websocket.Conn, error) {
	conn, err := websocket.Upgrade(w, r, responseHeader, 0, 0)
	if err != nil {
		return conn, err
	}
	return conn, nil
}
