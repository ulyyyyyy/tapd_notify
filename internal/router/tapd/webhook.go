package tapd

import (
	"github.com/gin-gonic/gin"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/ginresp"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
	"github.com/ulyyyyyy/tapd_notify/internal/model"
)

// Receive 接收Webhook数据
func Receive(c *gin.Context) {
	var body map[string]interface{}
	err := c.ShouldBindJSON(&body)
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

		}

		// 配置校验

		// TODO: 推送相关逻辑

	}()
	// 采用异步协程返回数据
	ginresp.NewSuccess(c, "OK")
}
