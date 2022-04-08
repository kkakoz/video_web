package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Conn struct {
	id          string
	conn        *websocket.Conn
	context     context.Context
	writeChan   chan any                                       // 写chan
	readFunc    func(msg []byte)                               // 读取之后处理
	errHandle   func(msg any, err error, conn *websocket.Conn) // 写err处理
	closeHandle func()                                         // close 处理
	cancel      context.CancelFunc
	once        sync.Once
}

type Options func(conn *Conn)

func CloseHandle(f func()) Options {
	return func(conn *Conn) {
		conn.closeHandle = f
	}
}

func ErrHandle(f func(msg any, err error, conn *websocket.Conn)) Options {
	return func(conn *Conn) {
		conn.errHandle = f
	}
}

func NewConn(ctx context.Context, id string, conn *websocket.Conn, readFunc func(msg []byte), opts ...Options) *Conn {
	ctx, cancel := context.WithCancel(ctx)
	c := &Conn{id: id, conn: conn, readFunc: readFunc, writeChan: make(chan any, 1),
		context: ctx, cancel: cancel}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (item *Conn) Writing() {
	for {
		select {
		case msg := <-item.writeChan:
			fmt.Println("in write chan case")
			err := item.conn.WriteMessage(1, []byte(fmt.Sprintf("%s", msg)))
			if err != nil {
				if item.errHandle != nil {
					item.errHandle(msg, err, item.conn)
				} else {
					log.Fatalln("write json err ", err)
				}
			}
		case <-item.context.Done():
			return
		}
	}
}

func (item *Conn) Reading() {
	for {
		select {
		case <-item.context.Done():
			return
		default:
			t, data, err := item.conn.ReadMessage()
			if err != nil {
				log.Println("reading err:", err)
				item.close()
			}
			log.Println("type = ", t)
			log.Println("data = ", string(data))
			item.readFunc(data)
		}
	}
}

func (item *Conn) Write(msg any) {
	item.writeChan <- msg
}

func (item *Conn) close() {
	item.once.Do(func() {
		item.closeHandle()
		item.cancel()
		item.conn.Close()
	})
}
