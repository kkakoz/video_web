package logic_test

import (
	"fmt"
	"github.com/kkakoz/ormx"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"testing"
	"time"
	"video_web/pkg/conf"
)

func Init() {
	conf.InitTestConfig()

	if _, err := ormx.New(conf.Conf()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}

	ormx.FlushDB()
}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}

func New(viper *viper.Viper) (*gorm.DB, error) {
	var err error
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", "3306")
	viper.SetDefault("db.name", "test")
	viper.SetDefault("db.console_log", "false")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?"+
		"charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.user"), viper.GetString("db.password"),
		viper.GetString("db.host"), viper.GetString("db.port"),
		viper.GetString("db.name"))
	loggerLevel := logger.Warn
	if viper.GetBool("db.log_info") { // 控制台打印普通sql
		loggerLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  loggerLevel, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	config := &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true, // AutoMigrate不会自动添加外键
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // table name prefix, table for `User` would be `t_users`
			SingularTable: true,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false, // skip the snake_casing of names
			//NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	}

	db, err := gorm.Open(mysql.Open(dns), config)
	return db, err
}
