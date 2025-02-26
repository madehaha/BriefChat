package router

import (
	"BriefChat/router/chater"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	Chater chater.ChaterRouterGroup
}

var RouterGroupApp = new(RouterGroup)

//全局化Router对象之后初始化

func (r RouterGroup) Init(rootRouterGroup *gin.RouterGroup) {
	r.Chater.Init(rootRouterGroup)
}
