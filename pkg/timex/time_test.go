package timex_test

import (
	"fmt"
	"testing"
	"time"
	"video_web/pkg/timex"
)

func TestTime(t *testing.T) {
	t1 := timex.Time{}
	def := time.Time{}
	if t1.Time == def {
		fmt.Println("t1 is null")
	}
	fmt.Println(t1.Unix())
}
