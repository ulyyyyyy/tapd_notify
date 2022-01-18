package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
	"strings"
	"time"
)

const (
	Local       = "local.yaml" // 本地开发
	Development = "dev.yaml"   // 测试服
	Production  = "prod.yaml"  // 正式服

	errFormat = "config file maybe not as expected (%s): %w"
	keyFormat = "config file maybe not as expected: %s"
)

// Active 当前使用的配置文件,
// 由 条件编译 `tags` 进行赋值
var Active string

var (
	DB           dbCfg
	GlobalSwtich globalSwitch
)

// dbCfg
type dbCfg struct {
	MySQL struct {
		DSN   string `structure:"dsn"`
		Conns struct {
			Max struct {
				Idle int `structure:"idle"`
				Open int `structure:"open"`
			} `structure:"max"`
		} `structure:"conns"`
	} `structure:"mysql"`

	Redis struct {
		Addr     string `structure:"addr"`
		Password string `structure:"password"`
		DB       int    `structure:"db"`
		Key      struct {
			Proxy string `structure:"proxy"`
		} `structure:"key"`
	} `structure:"redis"`

	Timeout struct {
		Connect time.Duration `structure:"connect"`
		Operate time.Duration `structure:"operate"`
	}
}

type globalSwitch struct {
	Proxy bool `structure:"proxy"`
}

// trimString
func trimString(key string, v interface{}) error {
	rVal := reflect.ValueOf(v)
	typ := rVal.Type()
	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < rVal.NumField(); i++ {
		if rVal.Field(i).Kind() == reflect.Struct {
			sKey := key + "." + typ.Field(i).Tag.Get("structure")
			if err := trimString(sKey, rVal.Field(i).Addr().Interface()); err != nil {
				return err
			}
		} else if rVal.Field(i).Kind() == reflect.String && rVal.Field(i).CanSet() {
			fKey := key + "." + typ.Field(i).Tag.Get("structure")
			s := strings.TrimSpace(rVal.Field(i).String())
			if len(s) == 0 {
				return fmt.Errorf(keyFormat, fKey)
			}
			rVal.Field(i).SetString(s)
		}
	}
	return nil
}

func VerifyAll() error {
	if err := verifyDB(); err != nil {
		return err
	}

	return nil
}

func verifyDB() error {
	key, val := "db", &DB
	if err := viper.UnmarshalKey(key, val); err != nil {
		return fmt.Errorf(errFormat, key, err)
	}

	if DB.MySQL.Conns.Max.Idle <= 0 {
		return fmt.Errorf(keyFormat, "db.mysql.conns.max.idle")
	}
	if DB.MySQL.Conns.Max.Open <= 0 {
		return fmt.Errorf(keyFormat, "db.mysql.conns.max.open")
	}

	if DB.Redis.DB < 0 {
		return fmt.Errorf(keyFormat, "db.redis.db")
	}
	return trimString(key, val)
}
