package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"video_web/pkg/gox"

	"github.com/spf13/viper"
)

type Application struct {
	name       string
	ctx        context.Context
	handler    http.Handler
	servers    []Server
	cancelFunc context.CancelFunc
}

func NewApplication(name string, handler http.Handler, servers []Server) *Application {
	ctx, cancelFunc := context.WithCancel(context.TODO())
	return &Application{name: name, handler: handler, servers: servers, ctx: ctx, cancelFunc: cancelFunc}
}

type Server interface {
	Start() error
	Stop() error
}

func (item *Application) Run() error {
	for _, serv := range item.servers {
		cur := serv
		gox.Go(func() {
			err := cur.Start()
			if err != nil {
				item.cancelFunc()
			}
		})
	}
	server := http.Server{
		Addr:    ":" + viper.GetString("app.port"),
		Handler: item.handler}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln("listen and server http err:", err)
		}
	}()
	go func() {
		quit := make(chan os.Signal, 1)
		<-quit
		time.Sleep(5 * time.Second)
		item.cancelFunc()
	}()

	<-item.ctx.Done()
	shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutDownCtx); err != nil {
		log.Fatalln("server shutdown err:", err)
	}
	for _, serv := range item.servers {
		err := serv.Stop()
		if err != nil {
			log.Fatalln("stop server err:", err)
		}
	}

	return nil
}
