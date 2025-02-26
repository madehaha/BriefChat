package service

import (
	"BriefChat/service/ai"
	"BriefChat/service/chat"
)

type ServiceGroup struct {
	Ai   ai.AiService
	Chat chat.ChatService
}

var ServiceGroupApp = new(ServiceGroup)
