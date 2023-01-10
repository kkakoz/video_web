package bootstrap

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"testing"
	"time"
	"video_web/pkg/conf"
	"video_web/pkg/ossx"
	"video_web/pkg/timex"
)

func TestUploadFiles(t *testing.T) {
	time.Local = time.FixedZone("UTC+8", 8*3600)
	pflag.Parse()
	conf.InitConfig("../configs/conf.yaml")

	err := initBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ossx.Init(viper.GetViper())
	if err != nil {
		t.Fatal(err)
	}

	name := "灵能百分百"
	brief := "平凡的中学二年级少年影山茂夫，因其微弱的存在感与名字茂夫的谐音被周遭人称为龙套（モブ），但不起眼的他其实是强大的天生超能力者。历经每一次的成长，龙套开始认为自己的超能力是危险的存在，为了不让超能力失控，龙套无意识的压抑著情感。虽然只想平凡的度过每一天，但各种麻烦却接二连三找上他，随着被压抑的情感在内心一点点膨胀，龙套体内积累的力量似乎也正蠢蠢欲动......"
	dir := "E:\\BaiduNetdiskDownload\\灵能百分百 第1季"
	var categoryId int64 = 1
	publishAt, err := timex.Parse("2006-01-02", "2016-07-12")
	if err != nil {
		t.Fatal(err)
	}
	err = uploadDirVideo(name, brief, dir, categoryId, &publishAt)
	if err != nil {
		t.Fatal(err)
	}
}
