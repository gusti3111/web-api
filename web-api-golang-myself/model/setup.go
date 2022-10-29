package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	dsn := "sqluser:password@tcp(127.0.0.1:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Input{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

}
