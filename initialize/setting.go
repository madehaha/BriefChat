package initialize

import (
	"BriefChat/global"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
)

type Config struct {
	App          global.AppSetting      `json:"app"`
	Database     global.DatabaseSetting `json:"database"`
	FileLocation global.FileSetting     `json:"file_location"`
	API          global.API             `json:"api"`
}

var config = &Config{}

func Setup() (*global.AppSetting, *global.DatabaseSetting, *global.FileSetting, *global.API) {
	fileContent, err := os.ReadFile("./conf.json")
	if err != nil {
		log.Fatalln("Can't read conf file")
	}
	err = jsoniter.Unmarshal(fileContent, config)
	if err != nil {
		log.Fatalln("Can't parse configuration")
	}

	return &config.App, &config.Database, &config.FileLocation, &config.API
}
