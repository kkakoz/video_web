package users_ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"video_web/pkg/errno"
)

type UserConns struct {
	users map[int64]*UserConn
}

var upgrade = websocket.Upgrader{}

func (item *UserConns) Add(w http.ResponseWriter, r *http.Request, userId int64) error {
	_, ok := item.users[userId]
	if ok {
		return errno.New400("已经连接")
	}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	go item.Reading(conn, userId)
	go item.Writing(conn)
	return nil
}

func (item *UserConns) Reading(conn *websocket.Conn, userId int64) {
	for {
		req := &Msg{}
		err := conn.ReadJSON(req)
		if err != nil {
			log.Fatalln(err)
		}
		switch req.Type {
		case 1:

		}
	}
}

func (item *UserConns) Writing(conn *websocket.Conn) {

}
