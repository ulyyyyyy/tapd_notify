package middleware

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
	"log"
)

// RocketConsumer rocket消费者
type RocketConsumer struct {
	C rocketmq.PushConsumer
}

func NewRocketConsumer(c rocketmq.PushConsumer) *RocketConsumer {
	return &RocketConsumer{C: c}
}

func (c *RocketConsumer) setConsumer(pushConsumer rocketmq.PushConsumer) {
	(*c).C = pushConsumer
}

var (
	con *RocketConsumer
)

// initPushConsumer 初始化并开启消费者
func initPushConsumer() (err error) {
	// 新建一个 PushConsumer，PushConsumer是被动接收型，即存在一条就接收
	C, err := rocketmq.NewPushConsumer(
		// 设置GroupName
		consumer.WithGroupName(groupName),
		//
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{address})),
		// 选择消费模式：集群消费，消费完其他人不能再读取
		consumer.WithConsumerModel(consumer.Clustering),
	)
	con = NewRocketConsumer(C)
	if err != nil {
		logger.Error(err.Error())
	}
	err = con.receive()
	if err != nil {
		logger.Error(err.Error())
	}
	return
}

func GetPushConsumerInstance() *RocketConsumer {
	if con == nil {
		_ = initPushConsumer()
	}
	return con
}

func (c *RocketConsumer) Subscribe(topicName string, ch chan string) int {
	err := c.C.Subscribe(topicName, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		// 一直接收数据
		for _, msg := range msgs {
			// 5秒数据
			if msg.ReconsumeTimes > 5 {
				log.Printf("msg ReconsumeTimes > 5,msg: %v", msg)
				return consumer.ConsumeSuccess, nil
			} else {
				log.Printf("subscribe callback: %v \n", msg)
				ch <- msg.String()
			}
		}
		return consumer.ConsumeRetryLater, nil
	})
	if err != nil {
		return 0
	}
	err = c.C.Start()
	if err != nil {
		log.Println(err.Error())
		c.ClosedConsumer()
	}
	return -1
}

// ClosedConsumer 关闭消费者
func (c *RocketConsumer) ClosedConsumer() {
	err := c.C.Shutdown()
	if err != nil {
		log.Printf("shundown Consumer error: %s", err.Error())
	}
}
