package configs

import (
	"fmt"
	"os"

	"github.com/NidzamuddinMuzakki/chat-golang-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_DB := os.Getenv("MYSQL_DATABASE")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	DATABASE_URI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DB)
	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.User{}, &models.Chat{})

	return nil
}
