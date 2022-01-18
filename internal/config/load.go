package config

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/ulyyyyyy/tapd_notify/configs"
)

func Load() error {
	viper.SetConfigType("yaml")
	return viper.ReadConfig(bytes.NewReader(configs.CfgFile))
}
