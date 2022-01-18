package proxy

import (
	"bytes"
	"github.com/ulyyyyyy/tapd_notify/configs"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/redis"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
	"io"
	"net/http"
)

// HttpClientSend 直接发送http请求数据
func HttpClientSend(body []byte) {
	// 获取redis中已经存入的转发目标地址
	proxies, err := redis.ListGetAll(configs.DB.Redis.Key.Proxy)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	for _, target := range proxies {
		target := target
		go func() {
			request, err := http.NewRequest(http.MethodPost, target, bytes.NewReader(body))
			if err != nil {
				logger.Error(err.Error())
				return
			}
			request.Header.Set("Content-Type", "application/json;charset=utf-8")
			client := &http.Client{}
			resp, err := client.Do(request)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			// 结束resp连接
			defer func() {
				_, _ = io.Copy(io.Discard, resp.Body)
				_ = resp.Body.Close()
			}()
		}()
	}
}
