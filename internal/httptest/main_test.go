package httptest__test

import (
	"encoding/json"
	"github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/handler"
	"github/kkakoz/video_web/internal/pkg/keys"
	"github/kkakoz/video_web/internal/router"
	"github/kkakoz/video_web/pkg/conf"
	"github/kkakoz/video_web/pkg/cryption"
	"github/kkakoz/video_web/pkg/mysqlx"
	"github/kkakoz/video_web/pkg/redis"
	"go.uber.org/fx"
	"log"
	"net/http"
	"testing"
	"time"
)

var server http.Handler

var userToken string

func TestMain(m *testing.M) {
	err := initServer()
	if err != nil {
		log.Fatalln("初始化失败:", err)
	}
	m.Run()
}

func initServer() error {
	conf.InitTestConfig()
	if _, err := mysqlx.New(viper.GetViper()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
	mysqlx.FlushDB()
	return fx.New(
		fx.NopLogger,
		handler.Provider,
		redis.Provider,
		router.Provider,
		fx.Supply(viper.GetViper()),
		fx.Populate(&server),
	).Err()
}

func loginUser() {
	client := redis.GetClient()
	user := &domain.User{
		ID:          10086,
		Name:        "test_user",
	}
	token := cryption.UUID()
	data, err := json.Marshal(user)
	convey.So(err, convey.ShouldBeNil)
	err = client.Set(keys.TokenKey(token), data, time.Minute).Err()
	convey.So(err, convey.ShouldBeNil)
	userToken = token
}

func logout() {
	userToken = ""
}