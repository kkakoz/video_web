package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Conn struct {
	conn      *websocket.Conn
	writeChan chan Msg
	readFunc  func(msg Msg)
}

func NewConn(conn *websocket.Conn, readFunc func(msg Msg)) *Conn {
	return &Conn{conn: conn, readFunc: readFunc, writeChan: make(chan Msg)}
}

type Msg struct {
	Type     uint   `json:"type"`
	TargetId uint   `json:"target_id"`
	Body     string `json:"body"`
}

func (item *Conn) Writing() {
	for {
		select {
		case msg := <-item.writeChan:
			fmt.Println("in write chan case")
			err := item.conn.WriteMessage(1, []byte(msg.Body))
			if err != nil {
				log.Fatalln("write json err ", err)
			}
		}
	}
}

func (item *Conn) Reading() {
	for {
		_, data, err := item.conn.ReadMessage()
		if err != nil {
			log.Println("reading err:", err)
		}
		log.Println("data = ", string(data))
		item.readFunc(Msg{
			Type:     0,
			TargetId: 0,
			Body:     string(data),
		})
		// req := &Msg{}
		// err := item.conn.ReadJSON(req)
		// if err != nil {
		// 	log.Println("reading err:", err)
		//
		// }
		// item.readFunc(*req)
	}
}

func (item *Conn) Write(msg Msg) {
	item.writeChan <- msg
}
