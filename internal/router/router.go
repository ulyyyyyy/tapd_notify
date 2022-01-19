package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ulyyyyyy/tapd_notify/internal/router/health"
	"github.com/ulyyyyyy/tapd_notify/internal/router/proxy"
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

	// 代理转发配置接口
	{
		router.POST("/proxy", proxy.InsertProxy)
		router.GET("/proxies", proxy.GetAllProxies)
		router.DELETE("/proxy", proxy.DeleteProxy)
	}
	{
		router.GET("/tapd/configs/:project", tapd.GetConfigByProject)
		router.GET("/tapd/config/:project/:id", tapd.GetConfigById)
		router.POST("/tapd/config/:project", tapd.CreateConfig)
		router.PUT("/tapd/config/:project/:id", tapd.UpdateConfigById)
		router.PUT("/tapd/config/status/:project/:id", tapd.UpdateStatusById)
		router.DELETE("/tapd/config/:project/:id", tapd.DeleteConfigById)
	}
	return
}
