package conf

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig(cfg string) {

	viper.SetConfigFile(cfg)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("read conf err:", err)
	}
}

func InitTestConfig() {
	s := "../../configs/test-conf.yaml"
	viper.SetConfigFile(s)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("read conf err:", err)
	}
}
