package ossx

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

type option struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Bucket          string `json:"bucket"`
}

var client = &oss.Client{}

var Bucket = &oss.Bucket{}

var Endpoint = ""

func Init(viper *viper.Viper) error {
	opt := &option{}
	err := viper.UnmarshalKey("oss", opt)
	if err != nil {
		return err
	}
	Endpoint = opt.Endpoint

	// 创建OSSClient实例。
	client, err = oss.New(Endpoint, opt.AccessKeyId, opt.AccessKeySecret)
	if err != nil {
		return err
	}
	// 获取存储空间。
	Bucket, err = client.Bucket(opt.Bucket)
	if err != nil {
		return err
	}
	// 上传文件。
	return nil
}
