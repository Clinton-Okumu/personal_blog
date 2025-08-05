package models

import "time"

type Article struct {
	BaseModel
	Title       string    `gorm:"not null"`
	Content     string    `gorm:"type:text;not null"`
	PublishedAt time.Time `gorm:"not null"`
}
