package initialize

import (
	"BriefChat/global"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	global.Global_APP_SETTING, global.Global_Db_SETTING, global.Global_FILE, global.Global_API = Setup()
	global.Global_Db = Gorm()
	Router(r)
}
