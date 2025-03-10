package ai

import (
	"BriefChat/api"
	"github.com/gin-gonic/gin"
)

type AIRouterGroup struct {
}

func (r *AIRouterGroup) Init(router *gin.RouterGroup) {
	aiRouter := router.Group("ai")
	//websocket no always use middleware but sec-websocket-protocol
	aiApi := api.ApiGroupApp.AiApiGroup
	{
		aiRouter.GET("/chatgpt", aiApi.HandleWebSocket)
	}
}
