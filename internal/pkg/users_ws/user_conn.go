package users_ws

import (
	"github.com/gorilla/websocket"
	"log"
)

type UserConn struct {
	userId    int64
	conn      *websocket.Conn
	writeChan chan Msg
}

type Msg struct {
	Type     uint   `json:"type"`
	TargetId uint   `json:"target_id"`
	Body     string `json:"body"`
}

const (
	TypeSendMsg = 1
	TypeSendImg = 2
)

func (item *UserConn) Writing() {
	for {
		select {
		case msg := <-item.writeChan:
			err := item.conn.WriteJSON(msg)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func (item *UserConn) Write(msg Msg) {
	item.writeChan <- msg
}
