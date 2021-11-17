package main

import (
	"github/kkakoz/video_web/pkg/conf"
	"github/kkakoz/video_web/pkg/mysqlx"
	"log"
)

func main() {
	viper, err := conf.GetConf("config/dev.conf")
	if err != nil {
		log.Fatalln("read conf err:", err)
	}
	_, err = mysqlx.New(viper)
	if err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
}
