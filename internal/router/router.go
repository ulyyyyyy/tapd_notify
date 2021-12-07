package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ulyyyyyy/tapd_notify/internal/router/health"
	"github.com/ulyyyyyy/tapd_notify/internal/router/tapd"
)

func InitRouter() (router *gin.Engine) {
	router = gin.New()

	{
		// 健康探针接口
		router.GET("/health", health.IsHealthy)
	}

	{
		// webhook消息接收接口
		router.POST("/webhook/tapd", tapd.Receive)
	}

	{
		// TODO 配置相关接口
	}
	return
}
