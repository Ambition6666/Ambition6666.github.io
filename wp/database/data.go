package data

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 打开数据库
func InitDB() {
	dsn := "host=localhost user=zty password=123456 dbname=zty port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
func GetDB() *gorm.DB {
	return DB
}
