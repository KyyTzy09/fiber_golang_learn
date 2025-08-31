package configs

import (
	"fiber/api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db * gorm.DB

func DbConnect() {
	dsn := "root@tcp(127.0.0.1:3306)/fiber_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
	Db = db
}
