package middleware

import (
	"github.com/spf13/viper"
)

const (
	_cfgKeyRocketMQ          = "rocketMQ"
	_cfgKeyRocketMQHost      = "rocketMQ.host"
	_cfgKeyRocketMQPort      = _cfgKeyRocketMQ + ".port"
	_cfgKeyRocketMQGroupName = _cfgKeyRocketMQ + ".groupName"
	_cfgKeyRocketMQTopicName = _cfgKeyRocketMQ + ".topicName"

	Retry int = 2
)

var (
	address   string
	groupName string
	topicName string
)

// Initialize 初始化mq相关配置
func Initialize() (err error) {
	// 载入相关配置
	address = viper.GetString(_cfgKeyRocketMQHost) + ":" + viper.GetString(_cfgKeyRocketMQPort) //"139.224.229.141:9876"
	groupName = viper.GetString(_cfgKeyRocketMQGroupName)
	topicName = viper.GetString(_cfgKeyRocketMQTopicName)
	// 初始化生产者
	err = initProducer(address, Retry)

	err = initPushConsumer()
	return
}
