package ws

import "go.uber.org/fx"

var Provider = fx.Provide(NewVideoConn, NewUserConn)
