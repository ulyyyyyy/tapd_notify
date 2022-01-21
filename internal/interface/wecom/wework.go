package wecom

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	_keyWecomDomain = "api.wework"
)

// GetAccessToken 获取access_token
func GetAccessToken(id, secret string) {
	weCompanyDomain := viper.GetString(_keyWecomDomain)
	_ = fmt.Sprint(weCompanyDomain, id, secret)
}
