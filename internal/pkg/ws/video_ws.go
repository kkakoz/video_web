package ws

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/kkakoz/pkg/gox"
	"github.com/kkakoz/pkg/redisx"
	"net/http"
	"sync"
	"video_web/internal/pkg/syncs"
	"video_web/pkg/errno"
)

type videoConn struct {
	videoMap syncs.Map[int64, *VideoConnLMap]
	redis    *redis.Client
}

type VideoConnLMap struct {
	syncs.Map[string, *Conn]
}

type VideoWsRes struct {
	Type    uint8 `json:"type"`
	Content any   `json:"content"`
}

const (
	VideoWsResTypeCount = 1
	VideoWsResTypeDanmu = 2
)

var videoConnOnce sync.Once
var _videoConn *videoConn

func VideoConn() *videoConn {
	videoConnOnce.Do(func() {
		client := redisx.Client()
		_videoConn = &videoConn{videoMap: syncs.NewMap[int64, *VideoConnLMap](), redis: client}
	})
	return _videoConn
}

func (item *videoConn) Add(w http.ResponseWriter, r *http.Request, videoId int64) error {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	newConn := NewConn(context.TODO(), r.RemoteAddr, conn, func(msg []byte) {
		item.sendBody(videoId, msg)
	},
		CloseHandle(func() { // 关闭处理
			item.videoMap.Do(func(m map[int64]*VideoConnLMap) {
				m[videoId].Del(r.RemoteAddr)
			})
		}),
		ErrHandle(func(msg any, err error, conn *websocket.Conn) { // 失败处理
			fmt.Println("msg = ", msg)
			fmt.Println("err = ", fmt.Sprintf("%+v", err))
		}))

	item.videoMap.Do(func(m map[int64]*VideoConnLMap) { // 加入到 video map中
		if m[videoId] == nil {
			m[videoId] = &VideoConnLMap{syncs.NewMap[string, *Conn]()}
		}
		if !m[videoId].SetFirst(r.RemoteAddr, newConn) {
			err = errno.New400("已经连接")
		}
	})
	if err != nil {
		return err
	}

	gox.Go(func() {
		newConn.Reading()
	})
	gox.Go(func() {
		newConn.Writing()
	})

	item.sendCount(videoId)
	return nil
}

func (item *videoConn) Send(videoId int64, res VideoWsRes) {
	connMap := item.videoMap.Get(videoId)
	connMap.Foreach(func(k string, conn *Conn) {
		conn.Write(res)
	})
}

func (item *videoConn) sendCount(videoId int64) {
	item.Send(videoId, VideoWsRes{
		Type:    VideoWsResTypeCount,
		Content: item.videoMap.Get(videoId).Len(),
	})
}

func (item *videoConn) sendBody(videoId int64, msg []byte) {
	item.Send(videoId, VideoWsRes{
		Type:    VideoWsResTypeCount,
		Content: string(msg),
	})
}
