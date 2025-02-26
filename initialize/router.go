package initialize

import (
	"BriefChat/global"
	"BriefChat/middleware"
	"BriefChat/router"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(middleware.Cors())

	rootRouterGroup := r.Group("")
	router.RouterGroupApp.Init(rootRouterGroup)

	r.Static("/public", *global.Global_FILE)
}
