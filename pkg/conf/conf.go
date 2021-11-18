package conf

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var cfg = pflag.StringP("config", "c", "configs/conf.yaml", "Configuration file.")

func init() {
	pflag.Parse()
}

func InitConfig() {
	viper.SetConfigFile(*cfg)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("read conf err:", err)
	}
}
