package logic_test

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"testing"
	request "video_web/internal/dto/request"
)

func TestUserLogic(t *testing.T) {

	err := userLogic.Register(context.TODO(), &request.RegisterReq{
		Name:     "lisi",
		Email:    "a@b.com",
		Password: "ttt",
	})
	if err != nil {
		t.Fatal(err)
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Fatal(err)
	}
	node.Generate().Int64()
	login, err := userLogic.Login(context.TODO(), &request.LoginReq{
		Name:     "a@b.com",
		Password: "ttt",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("token = ", login)
	// convey.Convey("user test", t, func(c convey.C) {
	// 	// c.Convey("login", func() {
	// 	// 	user, err := userLogic.GetUser(context.TODO(), 1)
	// 	// 	convey.ShouldBeNil(err)
	// 	// 	convey.ShouldNotBeNil(user)
	// 	// })
	// 	c.Convey("register", func() {
	//
	// 	})
	// })
}
