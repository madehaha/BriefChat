package chater

import (
	"BriefChat/global"
	"BriefChat/model/entity"
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
	token, info, err := chatService.Login(LoginReq)
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 绑定成功，继续处理
	e.JSON(http.StatusOK, gin.H{"token": token, "info": *(info)})

}

func (c *ChaterApiGroup) History(e *gin.Context) {
	account1, exists := e.Get("account")
	if !exists {
		// 如果上下文中没有 "username" 字段，返回错误
		e.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}
	records, err := chatService.History(account1.(string), e.Param("account"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	e.JSON(http.StatusOK, gin.H{"data": records})
}

func (c *ChaterApiGroup) AddFriend(e *gin.Context) {
	account1, exists := e.Get("account")
	if !exists {
		// 如果上下文中没有 "username" 字段，返回错误
		e.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}
	err := chatService.AddFriend(account1.(string), e.Param("account"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	e.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *ChaterApiGroup) AllFriends(e *gin.Context) {
	account1, exists := e.Get("account")
	if !exists {
		// 如果上下文中没有 "username" 字段，返回错误
		e.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}
	friends, err := chatService.AllFriend(account1.(string))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	e.JSON(http.StatusOK, gin.H{"data": friends})
}

func (c *ChaterApiGroup) GetSelf(e *gin.Context) {
	account1, exists := e.Get("account")
	if !exists {
		// 如果上下文中没有 "username" 字段，返回错误
		e.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}
	user, err := chatService.GetOneUser(account1.(string))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	e.JSON(http.StatusOK, gin.H{"data": *(user)})
}

func (c *ChaterApiGroup) GetOneUser(e *gin.Context) {
	account := e.Param("account")
	user, err := chatService.GetOneUser(account)
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	e.JSON(http.StatusOK, gin.H{"data": *(user)})
}

func (c *ChaterApiGroup) SendMessage(e *gin.Context) {
	var req request.MessageReq
	if err := e.ShouldBindJSON(&req); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account1, exists := e.Get("account")
	if !exists {
		// 如果上下文中没有 "username" 字段，返回错误
		e.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}
	if err := chatService.SendMessage(account1.(string), req.Receiver, req.Txt, req.Jpg); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	e.JSON(http.StatusOK, gin.H{"message": "send success"})
}

func (c *ChaterApiGroup) UploaderAvatar(e *gin.Context) {
	account1, exists := e.Get("account")
	if !exists {
		// 如果上下文中没有 "username" 字段，返回错误
		e.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}
	file, err := e.FormFile("avatar")
	//{注意前端是
	//headers: {
	//	"Content-Type": "multipart/form-data"
	//}
	//}
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"message": "上传失败"})
		return
	}
	path := account1.(string) + "_" + file.Filename
	filePath := *global.Global_FILE + "/" + path
	e.SaveUploadedFile(file, filePath)
	if err := chatService.Upload(account1.(string), entity.Info{
		Avatar: path,
		Name:   "",
	}); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"message": "头像上传失败"})
		return
	}
	e.JSON(http.StatusOK, gin.H{"avatar": "http://localhost:8080/public" + "/" + path})
}

func (c *ChaterApiGroup) UploaderInfo(e *gin.Context) {
	account1, exists := e.Get("account")
	if !exists {
		// 如果上下文中没有 "username" 字段，返回错误
		e.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}
	var req request.Info
	e.ShouldBindJSON(&req)
	if err := chatService.Upload(account1.(string), entity.Info{
		Avatar: "",
		Name:   req.Name,
	}); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"message": "头像上传失败"})
		return
	}
	e.JSON(http.StatusOK, gin.H{"message": "success"})

}
