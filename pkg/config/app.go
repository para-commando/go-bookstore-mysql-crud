package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Niveus@1233@tcp(localhost:3306)/BOOK_STORE?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

}

func GetDatabase() *gorm.DB {
	if DB == nil {
		ConnectDatabase()
	}
	return DB
}
