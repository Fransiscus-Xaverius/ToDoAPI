package migrations

import (
	"github.com/Fransiscus-Xaverius/ToDoAPI/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Todo{},
		&models.User{},
	)
}