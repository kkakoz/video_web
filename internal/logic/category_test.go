package logic_test

import (
	"context"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"video_web/internal/domain"
	"video_web/internal/logic/testdata"
)

func TestCategoryLogic(t *testing.T) {
	convey.Convey("category logic", t, func(c convey.C) {
		c.Convey("add", func() {
			for _, tt := range testdata.CategoryTests {
				err := categoryLogin.Add(context.TODO(), &domain.Category{Name: tt.In})
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