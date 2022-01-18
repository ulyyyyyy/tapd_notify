package tapd

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ulyyyyyy/tapd_notify/configs"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/ginresp"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
	"github.com/ulyyyyyy/tapd_notify/internal/model"
	"github.com/ulyyyyyy/tapd_notify/internal/proxy"
	"io"
	"sort"
)

var (
	allowedType = []string{""}
)

// Receive 接收Webhook数据
func Receive(c *gin.Context) {

	body, bsBody, err := getBody(c)
	logger.Info("Receive webhook.\t issueId: %s" + body["id"][12:])
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// 开关控制
	if configs.GlobalSwtich.Proxy {
		go func() {
			proxy.HttpClientSend(bsBody)
		}()
	}

	go func() {
		// 解析Request Body 中的数据
		if err != nil {
			logger.Error("[receive] json parse fail: " + err.Error())
			return
		}

		// 拉取相关配置
		configList, err := model.GetAllConfig()
		if err != nil {
			logger.Error("[config] get configList fail: " + err.Error())
			return
		}
		for _, config := range configList {
			// 检查是否符合条件
			fmt.Println(config)

		}

		// 配置校验

		// TODO: 推送相关逻辑

	}()
	// 采用异步协程返回数据
	ginresp.NewSuccess(c, nil)
}

// getBody 获取req的body数据。由于前置中间件可能bind了一次，所以需要判断 c.Get(gin.BodyBytesKey)中是否已经存在缓存数据
// 如果存在则直接取出数据返回，如果不存在则重新Bind一次返回
func getBody(c *gin.Context) (body map[string]string, bsBody []byte, err error) {
	// 由之前的 c.ShouldBindBodyWith 可以看到，方法将读到的 []byte 塞入进了 gin.BodyBytesKey 中，后续使用则可以直接获取。
	iBody, ok := c.Get(gin.BodyBytesKey)
	if ok {
		bsBody = iBody.([]byte)
	} else {
		bsBody, err = io.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error("" + err.Error())
			return
		}
	}
	if err = json.Unmarshal(bsBody, &body); err != nil {
		logger.Error("" + err.Error())
		return
	}
	return
}

// checkWebhookEvent 检查webhook事件是否需要被处理
func checkWebhookEvent(event string) bool {
	index := sort.SearchStrings(allowedType, event)
	// 检查是否包含在内，
	isExits := index <= len(allowedType) && index >= 0
	return isExits
}
