package snowx

import (
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"sync"
)

func GenerateInt64() int64 {
	return SnowFlake().Generate().Int64()
}

var once sync.Once
var node *snowflake.Node

func SnowFlake() *snowflake.Node {
	var err error
	once.Do(func() {
		node, err = snowflake.NewNode(viper.GetInt64("app.node"))
		if err != nil {
			node, _ = snowflake.NewNode(0)
		}
	})
	return node
}
