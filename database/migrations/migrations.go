package migrations

import (
	"github.com/meajudaaqui/user-service/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
