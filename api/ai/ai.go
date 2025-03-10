package ai

import (
	"BriefChat/global"
	"BriefChat/service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"sync"
)

type AiApiGroup struct {
	//Based Websocket
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

var aiService = service.ServiceGroupApp.Ai
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandleWebSocket 处理 WebSocket 连接
func (a *AiApiGroup) HandleWebSocket(c *gin.Context) {
	// 升级 HTTP 连接到 WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket 连接失败:", err)
		return
	}
	defer conn.Close()

	// 维护对话历史
	var chatHistory []Message
	var mutex sync.Mutex // 确保线程安全

	for {
		// 读取 WebSocket 消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket 读取失败:", err)
			break
		}

		userMessage := string(msg)
		fmt.Println("收到用户消息:", userMessage)

		// 添加用户消息到历史记录
		mutex.Lock()
		chatHistory = append(chatHistory, Message{
			Role:    "user",
			Content: userMessage,
		})
		mutex.Unlock()

		// 处理 OpenAI 聊天
		a.streamOpenAIResponse(conn, &chatHistory)
	}
}

// streamOpenAIResponse 以流式方式返回 OpenAI 回复
func (a *AiApiGroup) streamOpenAIResponse(conn *websocket.Conn, chatHistory *[]Message) {
	client := &http.Client{}
	// 发送给 OpenAI
	requestBody := RequestBody{
		Model: "qwen-plus", // 使用新的模型名称
		Messages: append([]Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
		}, *chatHistory...),
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("JSON 序列化失败:", err)
		return
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("请求创建失败:", err)
		return
	}

	// 设置请求头
	// 获取 API 密钥
	println(*global.Global_API)
	req.Header.Set("Authorization", "Bearer "+*global.Global_API)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("请求发送失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应失败:", err)
		return
	}

	// 将 AI 回复发送给前端
	err = conn.WriteMessage(websocket.TextMessage, bodyText)
	if err != nil {
		log.Println("WebSocket 发送失败:", err)
	}

	// 更新历史记录
	*chatHistory = append(*chatHistory, Message{
		Role:    "assistant",
		Content: string(bodyText),
	})
}
