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
		chatRouter.GET("all/:account", chatApi.History)
		chatRouter.POST("add/:account", chatApi.AddFriend)
		chatRouter.GET("friends", chatApi.AllFriends)
		chatRouter.POST("send", chatApi.SendMessage)
		chatRouter.GET("self", chatApi.GetSelf)
		chatRouter.GET("user/:account", chatApi.GetOneUser)
		chatRouter.POST("upload/avatar", chatApi.UploaderAvatar)
		chatRouter.POST("upload/info", chatApi.UploaderInfo)

	}
}
