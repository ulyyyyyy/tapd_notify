package middleware

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/ulyyyyyy/tapd_notify/internal/config"
	"log"
	"os"
	"testing"
)

func TestRocketConsumer_Subscribe(t *testing.T) {
	if err := config.Load(); err != nil {
		log.Printf("[configs] load failed: %s\n", err)
		os.Exit(1)
	}

	c := GetPushConsumerInstance()

	// 轮询遍历消息
	for {
		err := c.C.Subscribe(topicName, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			fmt.Printf("subscribe callback: %v \n", msgs)
			return consumer.ConsumeSuccess, nil
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
