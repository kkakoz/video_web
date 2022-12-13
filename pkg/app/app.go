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
	name    string
	handler http.Handler
	servers []Server
}

func NewApplication(name string, handler http.Handler, servers []Server) *Application {
	return &Application{name: name, handler: handler, servers: servers}
}

type Server interface {
	Start() error
	Stop() error
}

func (item *Application) Run(ctx context.Context) error {
	ctx, cancelFunc := context.WithCancel(context.TODO())
	for _, serv := range item.servers {
		cur := serv
		gox.Go(func() {
			err := cur.Start()
			if err != nil {
				cancelFunc()
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
		cancelFunc()
	}()

	<-ctx.Done()
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
