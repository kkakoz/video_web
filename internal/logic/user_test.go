package logic_test

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"testing"
	request "video_web/internal/dto/request"
)

func TestUserLogic(t *testing.T) {

	err := userLogic.Register(context.TODO(), &request.RegisterReq{
		Name:         "lisi",
		IdentityType: 1,
		Identifier:   "",
		Credential:   "",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = userLogic.Register(context.TODO(), &request.RegisterReq{
		Name:         "lisi",
		IdentityType: 1,
		Identifier:   "",
		Credential:   "",
	})
	if err != nil {
		t.Fatal(err)
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Fatal(err)
	}
	node.Generate().Int64()
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
