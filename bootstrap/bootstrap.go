package bootstrap

import (
	"github.com/spf13/pflag"
	"time"
	"video_web/pkg/conf"
)

var cfg = pflag.StringP("config", "c", "configs/conf.yaml", "Configuration file.")

func Run() error {
	time.Local = time.FixedZone("UTC+8", 8*3600)
	pflag.Parse()
	conf.InitConfig(*cfg)

	arg0 := pflag.Arg(0)

	switch arg0 {
	case "migrate":
		return migrate()
	case "job":
		return runJob()
	default:
		return run()
	}

}
