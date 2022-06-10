package snow_id

import (
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type idProduct struct {
	*snowflake.Node
}

func NewIdProduct(viper *viper.Viper) (*idProduct, error) {
	viper.SetDefault("app.node", 1)
	appNode := viper.GetInt64("app.node")
	node, err := snowflake.NewNode(appNode)
	if err != nil {
		return nil, err
	}
	return &idProduct{node}, nil
}

var Provider = fx.Provide(NewIdProduct)
