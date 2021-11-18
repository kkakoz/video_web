package server

import (
	"github.com/spf13/viper"
	"github/kkakoz/video_web/pkg/app"
	"go.uber.org/fx"
)

func Server(viper *viper.Viper) []app.Server {
	return []app.Server{}
}

var Provider = fx.Provide(Server)