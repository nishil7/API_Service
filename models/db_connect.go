package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Database_connect() {
	var err error
	dsn := "host=localhost user=nishil.bansal dbname=hospital port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed To connect to HMS Databse")
	}
	db.AutoMigrate(&Doctor{})
	db.AutoMigrate(&Patient{})
	DB = db

}
