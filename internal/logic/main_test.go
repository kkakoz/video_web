package logic_test

import (
	"context"
	"github.com/kkakoz/ormx"
	"log"
	"testing"
	"video_web/internal/model"
	"video_web/pkg/conf"
)

func Init() {
	conf.InitTestConfig()
	if _, err := ormx.New(conf.Conf()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
	ormx.FlushDB()
	ormx.DB(context.TODO()).AutoMigrate(&model.User{})
}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}
