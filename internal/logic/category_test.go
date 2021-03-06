package logic_test

import (
	"context"
	"sync"
	"testing"
	"video_web/internal/logic/testdata"
	"video_web/internal/model"

	"github.com/smartystreets/goconvey/convey"
)

func TestCategoryLogic(t *testing.T) {
	group := sync.WaitGroup{}
	group.Wait()
	convey.Convey("category logic", t, func(c convey.C) {
		c.Convey("add", func() {
			for _, tt := range testdata.CategoryTests {
				err := categoryLogin.Add(context.TODO(), &model.Category{Name: tt.In})
				convey.ShouldEqual(err, tt.Expected)
			}
		})
		c.Convey("list", func() {
			list, err := categoryLogin.List(context.TODO())
			convey.ShouldBeNil(err)
			convey.ShouldEqual(len(list), 1)
		})
	})
}
