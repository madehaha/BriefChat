package global

import (
	"gorm.io/gorm"
)

var (
	Global_Db          *gorm.DB
	Global_APP_SETTING *AppSetting
	Global_Db_SETTING  *DatabaseSetting
	Global_FILE        *FileSetting
	Global_API         *API
	//Global_AI_CLIENT   *openai.Client
)

const (
	JWT_TOKEN_PREFIX = "Bearer "
)

type AppSetting struct {
	Domain    string `json:"domain"`
	Address   string `json:"address"`
	JwtSecret string `json:"jwt_secret"`
}

type DatabaseSetting struct {
	Type     string `json:"type"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	//TablePrefix string `json:"table_prefix"`
}

type FileSetting = string

type API = string
