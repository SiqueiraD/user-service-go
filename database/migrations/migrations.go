package migrations

import (
	"github.com/siqueirad/user-service-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
