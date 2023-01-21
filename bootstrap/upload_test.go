package bootstrap

import (
	"github.com/kkakoz/pkg/conf"
	"testing"
	"time"
	"video_web/pkg/ossx"
	"video_web/pkg/timex"
)

func TestUploadFiles(t *testing.T) {
	time.Local = time.FixedZone("UTC+8", 8*3600)

	err := initBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ossx.Init(conf.Conf())
	if err != nil {
		t.Fatal(err)
	}

	name := "冰菓"
	brief := "在众多将要展开「玫瑰色」生活的高中生之中，本作的男主角折木奉太郎却是一个「灰色」的节能主义者。凡是没必要的事就不做，因为不想后悔，被人说是疏离、厌世也无所谓，因为这就是他的作风。这样的折木奉太郎，却因为姐姐的命令而进入了濒临废社的「古籍研究社」。\n研究社虽然好不容易招到了四名新社员，但却又卷入了四十五年前社长突然肄业的谜团之中。社长当年留下的名为「冰菓」的社刊，内里究竟隐藏了什么神秘的讯息呢……"
	dir := "E:\\BaiduNetdiskDownload\\冰菓"
	var categoryId int64 = 2
	publishAt, err := timex.Parse("2006-01-02", "1975-07-12")
	if err != nil {
		t.Fatal(err)
	}
	err = uploadDirVideo(name, brief, dir, categoryId, &publishAt)
	if err != nil {
		t.Fatal(err)
	}
}
