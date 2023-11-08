package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	Name      string
	Blogs     []Blog `gorm:"many2many:blog_categories"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
