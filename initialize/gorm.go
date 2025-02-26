package initialize

import (
	"BriefChat/global"
	"BriefChat/model/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func Gorm() *gorm.DB {
	var db *gorm.DB
	var err error
	databaseSetting := global.Global_Db_SETTING
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", databaseSetting.User, databaseSetting.Password, databaseSetting.Host, databaseSetting.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移数据库，创建表
	db.AutoMigrate(&entity.Chat{})
	return db
}
