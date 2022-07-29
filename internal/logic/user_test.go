package logic_test

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"testing"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/pkg/redisx"
)

func TestUserLogic(t *testing.T) {

	err := logic.User().Register(context.TODO(), &dto.Register{
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
	login, err := logic.User().Login(context.TODO(), &dto.Login{
		Name:     "a@b.com",
		Password: "ttt",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("token = ", login)
}

func TestA(t *testing.T) {
	result, err := redisx.Client().Get("test1:sss").Result()
	fmt.Println(result, err)
}
