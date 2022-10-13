package logic_test

import (
	"fmt"
	"testing"
	"video_web/pkg/mathx"
)

func TestCategoryLogic(t *testing.T) {

	//convey.Convey("_category logic", t, func(c convey.C) {
	//	c.Convey("add", func() {
	//		for _, tt := range testdata.CategoryTests {
	//			err := logic.Category().Add(context.TODO(), &entity.Category{Name: tt.In})
	//			convey.ShouldEqual(err, tt.Expected)
	//		}
	//	})
	//	c.Convey("list", func() {
	//		list, err := logic.Category().List(context.TODO())
	//		convey.ShouldBeNil(err)
	//		convey.ShouldEqual(len(list), 1)
	//	})
	//})
	//
	fmt.Println(mathx.Sum([]uint{1, 3, 4, 5, 5}))
}
