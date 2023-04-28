package configs

import (
	"github.com/NidzamuddinMuzakki/chat-golang-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:123@tcp(localhost:3306)/test-api?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

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
