package chater

import (
	"BriefChat/model/request"
	"BriefChat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var chatService = service.ServiceGroupApp.Chat

type ChaterApiGroup struct {
}

func (c *ChaterApiGroup) Register(e *gin.Context) {
	var RegisterReq request.Register

	if err := e.ShouldBindJSON(&RegisterReq); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := chatService.Register(RegisterReq); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// 绑定成功，继续处理
	e.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}

func (c *ChaterApiGroup) Login(e *gin.Context) {
	var LoginReq request.Login

	if err := e.ShouldBindJSON(&LoginReq); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := chatService.Login(LoginReq)
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 绑定成功，继续处理
	e.JSON(http.StatusOK, gin.H{"token": token})

}

func (c *ChaterApiGroup) History(e *gin.Context) {
	account1, exists := e.Get("account")
	if !exists {
		// 如果上下文中没有 "username" 字段，返回错误
		e.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}
	records, err := chatService.History(account1.(string), e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	e.JSON(http.StatusOK, gin.H{"data": records})
}
