package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ncutbbs/config"
	"os"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.DBUser, config.Config.DBPassword, config.Config.DBHost, config.Config.DBPort, config.Config.DBName)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(0)
	} else {
		DB = db
	}
}
