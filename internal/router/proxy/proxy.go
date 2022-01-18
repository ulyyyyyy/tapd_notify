package proxy

import (
	"github.com/gin-gonic/gin"
	"github.com/ulyyyyyy/tapd_notify/configs"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/ginresp"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/redis"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
	"strconv"
)

// GetAllProxies 获取所有Proxies数据
func GetAllProxies(c *gin.Context) {
	keys, err := redis.ListGetAll(configs.DB.Redis.Key.Proxy)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrProxyNotExits, nil)
		return
	}

	ginresp.NewSuccess(c, keys)
}

func DeleteProxy(c *gin.Context) {

	var reqBody proxy
	err := c.BindJSON(&reqBody)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrReceiveBody, reqBody)
		return
	}
	rem, err := redis.LRem(configs.DB.Redis.Key.Proxy, 1, reqBody.Target)
	logger.Info("remove " + strconv.FormatInt(rem, 10) + " value(s)")
	ginresp.NewSuccess(c, reqBody.Target)
}

func InsertProxy(c *gin.Context) {
	var reqBody proxy
	err := c.BindJSON(&reqBody)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrReceiveBody, reqBody)
		return
	}

	if err = redis.RPush(configs.DB.Redis.Key.Proxy, reqBody.Target); err != nil {
		ginresp.NewFailure(c, ginresp.ErrReceiveBody, reqBody)
		return
	}

	ginresp.NewSuccess(c, nil)
}

type proxy struct {
	Target string `json:"target"`
}
