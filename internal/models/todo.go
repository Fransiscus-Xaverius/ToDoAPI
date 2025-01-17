package models

import (
    "gorm.io/gorm"
    "time"
)

type Todo struct {
    gorm.Model
    ID          uint       
    UserID      string       
    Title       string     
    Description string     
    Completed   bool       
    CompletedAt *time.Time 
}
