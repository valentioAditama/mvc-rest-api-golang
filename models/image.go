package models

import (
	"gorm.io/gorm"
	"time"
)

type Image struct {
	gorm.Model
	path      string
	Blogs     []Blog `gorm:"many2many:blog_images"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
