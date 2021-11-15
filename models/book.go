package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint64         `json:"_id" gorm:"primayKey"`
	Title       string         `json:"title" gorm:"title"`
	Description string         `json:"description" gorm:"description"`
	Price       float64        `json:"price" gorm:"price"`
	Author      string         `json:"author" gorm:"author"`
	ImageURL    string         `json:"img_url" gorm:"img_url"`
	CreatedAt   time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
}
