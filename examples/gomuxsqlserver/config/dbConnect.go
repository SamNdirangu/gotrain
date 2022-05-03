package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const dsn = "sam:123@tcp(127.0.0.1:3308)/gotest?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() {
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
