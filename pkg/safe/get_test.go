package safe_test

import (
	"fmt"
	"testing"
	"video_web/pkg/safe"
)

func TestSafeGet(t *testing.T) {

	var a []string

	fmt.Println(safe.Get(func() string {
		return a[1]
	}))

	fmt.Println(safe.GetDef(func() string {
		return a[1]
	}, "safe"))
}
