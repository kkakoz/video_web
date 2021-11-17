package app

import (
	"net"
	"net/http"
)

type Application struct {
	handler http.Handler
	servers []Server
	Name    string
	Port    string
}

func NewApplication(handler http.Handler, servers []Server) *Application {
	return &Application{handler: handler, servers: servers}
}

type Server interface {
	Start() error
	Stop() error
}

func (item *Application) Run() error {
	listen, err := net.Listen("tcp", ":" + item.Port)
	if err != nil {
		return err
	}
	go func() {

	}()
	err = http.Serve(listen, item.handler)
	return nil
}

func (Application) Stop() {

}
