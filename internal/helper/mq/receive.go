package middleware

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
)

// receive 接受消息队列消息
func (c *RocketConsumer) receive() (err error) {
	err = c.C.Start()
	// 开启消费者
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = c.C.Subscribe(topicName, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			logger.Info(string(msg.Body))
			// TODO push 业务
			//errPush := message.Push(msg)
			//if errPush != nil {
			//	// TODO: 鉴别是什么问题
			//	logger.Error(errPush.Error())
			//}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		logger.Error(err.Error())
	}
	return
}
