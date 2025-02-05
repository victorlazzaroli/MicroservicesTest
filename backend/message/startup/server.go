package startup

import (
	"github.com/victorlazzaroli/microservicesTest/message/api/message/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = database.AutoMigrate(&model.Message{})

	if err != nil {
		return
	}

	DB = database
}
