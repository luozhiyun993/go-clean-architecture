package app

import (
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//dsn := "root:123456@tcp(127.0.0.1:3306)?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	return &gorm.DB{}
}
