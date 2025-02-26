package api

import (
	"BriefChat/api/ai"
	"BriefChat/api/chater"
)

type ApiGroup struct {
	ChatApiGroup chater.ChaterApiGroup
	AiApiGroup   ai.AiApiGroup
}

var ApiGroupApp = new(ApiGroup)
