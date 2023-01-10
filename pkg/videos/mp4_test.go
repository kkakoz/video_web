package videos

import (
	"fmt"
	"os"
	"testing"
)

func TestGetMP4Duration(t *testing.T) {
	file, err := os.Open("E:\\BaiduNetdiskDownload\\灵能百分百 第1季\\第1集.mp4")
	if err != nil {
		panic(err)
	}
	duration, err := GetMP4Duration(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(duration)
}
