package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	config := LoadConfig()

	var err error
	DB, err = gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		fmt.Printf("error while connecting to database: %v", err.Error())
		panic("failed to connect to database")
	}
}

func GetDatabase() *gorm.DB {
	if DB == nil {
		ConnectDatabase()
	}
	return DB
}
