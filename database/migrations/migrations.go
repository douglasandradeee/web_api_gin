package migrations

import (
	"github.com/douglasandradeee/web_api_gin/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
