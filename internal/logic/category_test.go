package logic_test

import (
	"context"
	"golang.org/x/exp/constraints"
	"testing"
	"video_web/internal/logic"
	"video_web/internal/logic/testdata"
	"video_web/internal/model/entity"

	"github.com/smartystreets/goconvey/convey"
)

func TestCategoryLogic(t *testing.T) {

	convey.Convey("_category logic", t, func(c convey.C) {
		c.Convey("add", func() {
			for _, tt := range testdata.CategoryTests {
				err := logic.Category().Add(context.TODO(), &entity.Category{Name: tt.In})
				convey.ShouldEqual(err, tt.Expected)
			}
		})
		c.Convey("list", func() {
			list, err := logic.Category().List(context.TODO())
			convey.ShouldBeNil(err)
			convey.ShouldEqual(len(list), 1)
		})
	})
}

func SliceMap[T any, K constraints.Ordered, V any](s []T, f func(v T) (K, V)) map[K]V {

	m := map[K]V{}
	for _, v := range s {
		k, v := f(v)
		m[k] = v
	}
	return m
}
