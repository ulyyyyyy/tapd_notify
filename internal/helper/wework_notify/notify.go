// Package wework_notify 维护企业微信消息推送相关类，接入开源项目 WeCom-Bot-API
package wework_notify

import (
	botApi "github.com/electricbubble/wecom-bot-api"
	"log"
)

func Push(message string, receiverList []string) {
	_ = botApi.NewWeComBot("test")
}

func PushSingle(message string, receiver string) {
	bot := botApi.NewWeComBot(receiver)
	if err := bot.PushMarkdownMessage(message); err != nil {
		log.Fatalln(err)
	}
}

func PushMerge(receiver string, mergeNum int) {

}
