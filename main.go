package main

import (
	"BriefChat/global"
	"BriefChat/initialize"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 初始化Gin引擎
	r := gin.Default()

	initialize.Init(r)
	// 路由配置

	// 启动服务器
	if err := r.Run(global.Global_APP_SETTING.Address); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
