// Package model
package model

import (
	"github.com/ulyyyyyy/tapd_notify/internal/helper/mysql"
	"gorm.io/gorm/schema"
)

const (
	_tableName = "tapd_config"
)

var (
	_ schema.Tabler = (*Config)(nil)
	// _ callbacks.BeforeCreateInterface = (*JobRecord)(nil)
)

type Config struct {
	Id          int    // 配置Id
	Name        string // 配置名称
	Description string // 配置描述

	// ReceiveType StrList // 接收类型，指 task:create / task:update 等
	ConditionType int            // 配置条件类型
	Condition     JsonListParser // 配置条件

	SummaryField StrList // 推送内容
	PushList     StrList // 推送人
}

func (cfg *Config) TableName() string {
	return _tableName
}

func (cfg Config) Insert() error {
	return mysql.Insert(cfg, mysql.DB())
}

func (cfg Config) Delete() error {
	return mysql.Delete(cfg, mysql.DB())
}

// GetAllConfig 获取所有数据
func GetAllConfig() (configList []Config, err error) {
	err = mysql.DB().Find(&configList).Error
	return configList, err
}
