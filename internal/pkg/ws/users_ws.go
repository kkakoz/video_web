package ws

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/kkakoz/pkg/gox"
	"net/http"
	"video_web/internal/pkg/syncs"
	"video_web/pkg/errno"
)

type UserConn struct {
	userMap syncs.Map[int64, *Conn]
	redis   redis.Client
}

type UserWsRes struct {
	FromType int64
	FromId   int64
}

func NewUserConn(redis redis.Client) *UserConn {
	return &UserConn{userMap: syncs.NewMap[int64, *Conn](), redis: redis}
}

func (item *UserConn) Add(w http.ResponseWriter, r *http.Request, userId int64) error {
	item.redis.Do("bf.add")
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	newConn := NewConn(context.TODO(), fmt.Sprintf("%d", userId), conn, func(msg []byte) {
		item.handler(msg)
	}, CloseHandle(func() {
		item.userMap.Del(userId)
	}), ErrHandle(func(msg any, err error, conn *websocket.Conn) {
		fmt.Println("msg = ", msg)
		fmt.Println("err = ", fmt.Sprintf("%+v", err))
	}))
	if item.userMap.SetFirst(userId, newConn) {
		return errno.New400("已经连接")
	}
	gox.Go(func() {
		newConn.Reading()
	})
	gox.Go(func() {
		newConn.Writing()
	})
	return nil
}

func (item *UserConn) handler(msg []byte) {

}

func (item *UserConn) Send(userId int64, res UserWsRes) bool {
	conn := item.userMap.Get(userId)
	if conn == nil {
		return false
	}
	conn.Write(res)
	return true
}
