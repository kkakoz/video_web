package conf

import (
	"github.com/spf13/viper"
	"os"
)

func GetConf(filepath string) (*viper.Viper, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	err = viper.ReadConfig(file)
	return viper.GetViper(), err
}
