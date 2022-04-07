package ws

import (
	"github.com/go-redis/redis"
	"net/http"
	"video_web/internal/pkg/syncs"
	"video_web/pkg/errno"
	"video_web/pkg/gox"
)

type UserConn struct {
	userMap syncs.Map[int64, *Conn]
	redis   redis.Client
}

func NewUserConn(redis redis.Client) *UserConn {
	return &UserConn{userMap: syncs.NewMap[int64, *Conn](), redis: redis}
}

func (item *UserConn) Add(w http.ResponseWriter, r *http.Request, userId int64) error {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	newConn := NewConn(conn, func(msg Msg) {
		item.handler(msg)
	})
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

func (item *UserConn) handler(msg Msg) {

}
