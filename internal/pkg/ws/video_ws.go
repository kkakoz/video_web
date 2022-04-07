package ws

import (
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"video_web/internal/pkg/syncs"
	"video_web/pkg/gox"
)

type VideoConn struct {
	videoMap syncs.Map[int64, []*Conn]
	redis    *redis.Client
}

func NewVideoConn(redis *redis.Client) *VideoConn {
	return &VideoConn{videoMap: syncs.NewMap[int64, []*Conn](), redis: redis}
}

func (item *VideoConn) Add(w http.ResponseWriter, r *http.Request, videoId int64) error {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	newConn := NewConn(conn, func(msg Msg) {
		item.sendBody(videoId)
	})

	item.videoMap.Do(func(m map[int64][]*Conn) {
		m[videoId] = append(m[videoId], newConn)
	})

	gox.Go(func() {
		newConn.Reading()
	})
	gox.Go(func() {
		newConn.Writing()
	})

	item.sendCount(videoId)
	return nil
}

func (item *VideoConn) handler(msg Msg) {
}

func (item *VideoConn) sendCount(videoId int64) {
	connList := item.videoMap.Get(videoId)
	for _, conn := range connList {
		conn.Write(Msg{
			Type:     0,
			TargetId: 0,
			Body:     fmt.Sprintf("%d", len(connList)),
		})
	}
}

func (item *VideoConn) sendBody(videoId int64) {
	connList := item.videoMap.Get(videoId)
	for _, conn := range connList {
		conn.Write(Msg{
			Type:     0,
			TargetId: 0,
			Body:     "hello",
		})
	}
}
