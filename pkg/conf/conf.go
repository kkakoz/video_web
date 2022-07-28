package conf

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var cfg = pflag.StringP("config", "c", "configs/conf.yaml", "Configuration file.")

var conf *viper.Viper

func InitConfig() {
	pflag.Parse()
	viper.SetConfigFile(*cfg)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("read conf err:", err)
	}
	conf = viper.GetViper()
}

func InitTestConfig() {
	s := "test-conf.yaml"
	viper.SetConfigFile(s)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("read conf err:", err)
	}
	conf = viper.GetViper()
}

func Conf() *viper.Viper {
	return conf
}
