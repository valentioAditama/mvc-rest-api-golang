package models

import (
	"time"
)

type Blog struct {
	ID         uint   `gorm:"primarykey"`
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	categories []Category `gorm:"many2many:blog_categories"`
	images     []Image    `gorm:"many2many:blog_images"`
	users      []User     `gorm:"many2many:blog_users"`
}
