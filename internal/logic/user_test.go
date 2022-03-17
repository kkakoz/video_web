package logic_test

import (
	"context"
	"testing"
	"video_web/internal/dto/request"
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
