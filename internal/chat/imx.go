package imx

import (
	"context"
	"fmt"
	"github.com/kkakoz/gim"
	"github.com/kkakoz/pkg/logger"
	"github.com/pkg/errors"
	"time"
)

type Chat struct {
}

func (c *Chat) Start(ctx context.Context) error {
	var srv gim.Server

	handler := &ServerHandler{}

	srv.SetReadWait(time.Minute)
	srv.SetAcceptor(handler)
	srv.SetMessageListener(handler)
	srv.SetStateListener(handler)
	srv.SetChannelMap(gim.NewChannelMap())

	err := srv.Start()
	if err != nil {
		logger.Fatal("server start error: " + err.Error())
	}
	return nil
}

func (c *Chat) Stop(ctx context.Context) error {

	return nil
}

// ServerHandler ServerHandler
type ServerHandler struct {
}

// Accept this connection
func (h *ServerHandler) Accept(conn gim.Conn, timeout time.Duration) (string, error) {
	// 1. 读取：客户端发送的鉴权数据包
	frame, err := conn.ReadFrame()
	if err != nil {
		return "", err
	}
	// 2. 解析：数据包内容就是userId
	userID := string(frame.GetPayload())
	// 3. 鉴权：这里只是为了示例做一个fake验证，非空
	if userID == "" {
		return "", errors.New("user id is invalid")
	}
	return userID, nil
}

// Receive default listener
func (h *ServerHandler) Receive(ag gim.Agent, payload []byte) {
	ack := string(payload) + " from server "
	_ = ag.Push([]byte(ack))
}

// Disconnect default listener
func (h *ServerHandler) Disconnect(id string) error {
	logger.Warn(fmt.Sprintf("disconnect %s", id))
	return nil
}
