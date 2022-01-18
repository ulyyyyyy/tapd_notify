package middleware

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
)

// RocketProducer 生产者
type RocketProducer struct {
	P rocketmq.Producer
}

func NewRocketProducer(p rocketmq.Producer) *RocketProducer {
	return &RocketProducer{P: p}
}

var (
	pro *RocketProducer
)

// initProducer 初始化生产者
func initProducer(address string, retry int) (err error) {
	addr, err := primitive.NewNamesrvAddr(address)
	if err != nil {
		return err
	}
	p, err := rocketmq.NewProducer(
		//
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{address})),
		//
		producer.WithRetry(retry),
		// 群组名
		producer.WithGroupName(groupName),
		// 服务地址
		producer.WithNameServer(addr),
	)
	pro = NewRocketProducer(p)
	err = pro.P.Start()
	if err != nil {
		logger.Error(err.Error())
	}
	return
}

// getProducerInstance 获取生产者单例
func getProducerInstance() *RocketProducer {
	if pro == nil {
		err := initProducer(address, Retry)
		if err != nil {
			logger.Error(err.Error())
		}
	}
	return pro
}

// SendSync 同步模式发送消息
func (p *RocketProducer) SendSync(message []byte) (err error) {
	_, err = p.P.SendSync(context.Background(), primitive.NewMessage(topicName, message))
	if err != nil {
		logger.Error(err.Error())
	}
	return
}

func SendSync(message []byte) (err error) {
	p := getProducerInstance()
	return p.SendSync(message)
}
