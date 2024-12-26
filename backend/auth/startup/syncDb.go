package startup

import (
	"log"

	"github.com/victorlazzaroli/microservicesTest/auth/models"
	"gorm.io/gorm"
)

func SyncDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
