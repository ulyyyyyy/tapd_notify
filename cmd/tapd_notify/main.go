package main

import (
	"fmt"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
	"github.com/ulyyyyyy/tapd_notify/internal/router"
	"log"
	"net/http"
	"os"
)

func main() {
	// 用以Tapd Webhook 消息推送功能
	log.Println("start...")

	server := http.Server{
		Addr:    ":9110",
		Handler: router.InitRouter(),
	}
	// 初始化方法
	initialize()

	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}

// initialize 初始化方法
func initialize() {
	var err error

	// 初始化 logger 日志打印组件，如果失败则退出
	err = logger.Initialize()
	if err != nil {
		fmt.Printf("[logger]init fail: %s", err.Error())
		os.Exit(1)
	}

	//
}
