package models

import (
    "gorm.io/gorm"
    "time"
)

type Todo struct {
    ID          uint       `gorm:"primaryKey"`
    Title       string     `gorm:"size:255;not null" json:"title"`
    Description string     `gorm:"size:255;not null" json:"description"`
    Completed   bool       `gorm:"default:false" json:"completed"`
    CompletedAt *time.Time `gorm:"default:null" json:"completed_at"`
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&Todo{})
}
