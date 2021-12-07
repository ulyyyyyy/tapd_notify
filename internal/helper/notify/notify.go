// Package notify 维护企业微信消息推送相关类，接入开源项目 WeCom-Bot-API
package notify

import (
	botApi "github.com/electricbubble/wecom-bot-api"
)

type notifyBot struct {
	webhook string
	key     string
}

var (
	_   botApi.WeComBot = (*notifyBot)(nil)
	bot botApi.WeComBot
)

func (n notifyBot) PushTextMessage(content string, opts ...botApi.TextMsgOption) error {
	panic("implement me")
}

func (n notifyBot) PushMarkdownMessage(content string) error {
	panic("implement me")
}

func (n notifyBot) PushImageMessage(img []byte) error {
	panic("implement me")
}

func (n notifyBot) PushNewsMessage(art botApi.Article, articles ...botApi.Article) error {
	panic("implement me")
}

func (n notifyBot) PushFileMessage(media botApi.Media) error {
	panic("implement me")
}

func (n notifyBot) PushTemplateCardTextNotice(mainTitle botApi.TemplateCardMainTitleOption, cardAction botApi.TemplateCardAction, opts ...botApi.TemplateCardOption) error {
	panic("implement me")
}

func (n notifyBot) PushTemplateCardNewsNotice(mainTitle botApi.TemplateCardMainTitleOption, cardImage botApi.TemplateCardImageOption, cardAction botApi.TemplateCardAction, opts ...botApi.TemplateCardOption) error {
	panic("implement me")
}

//func newWeComBot(webhook string, key string) botApi.WeComBot {
//	bot := new(notifyBot)
//	bot.webhook = webhook
//	bot.key = key
//	return bot
//}

func (n notifyBot) UploadFile(filename string) (media botApi.Media, err error) {
	panic("implement me")
}

func initBot() (bot botApi.WeComBot) {
	return
}

// Push
func Push(message string, receiverList []string) {
	_ = botApi.NewWeComBot("test")
}
