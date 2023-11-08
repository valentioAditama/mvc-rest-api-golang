package models

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	gorm.Model
	Title      string
	Content    string
	CreateAt   time.Time
	UpdateAt   time.Time
	categories []Category `gorm:"many2many:blog_categories"`
	images     []Image    `gorm:"many2many:blog_images"`
	users      []User     `gorm:"many2many:blog_users"`
}
