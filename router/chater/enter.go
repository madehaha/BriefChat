package chater

import (
	"BriefChat/api"
	"BriefChat/middleware"
	"github.com/gin-gonic/gin"
)

type ChaterRouterGroup struct {
}

func (c *ChaterRouterGroup) Init(router *gin.RouterGroup) {

	chatRouterWithoutJwt := router.Group("chat")
	chatRouter := router.Group("chat").Use(middleware.JwtMiddleWare())
	chatApi := api.ApiGroupApp.ChatApiGroup
	{
		chatRouterWithoutJwt.POST("register", chatApi.Register)
		chatRouterWithoutJwt.POST("login", chatApi.Login)
		chatRouter.GET("all/:id", chatApi.History)
		//chatRouter.POST("send", chatApi.Send)
		//chatRouter.GET("self", chatApi.ChaterInfo)
		//chatRouter.POST("change", chatApi.ChangeInfo)

	}
}
