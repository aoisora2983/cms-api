package db

import (
	"cms/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

func GetDB() *gorm.DB {
	if database == nil {
		database = Init()
	}

	return database
}

func Init() *gorm.DB {
	host := config.DBHost()
	user := config.DBUser()
	pass := config.DBPass()
	name := config.DBName()
	port := config.DBPort()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if config.IsLocal() {
		db.Logger = db.Logger.LogMode(logger.Info)
	}

	return db
}
