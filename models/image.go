package models

import "time"

type Image struct {
	ID        uint `gorm:"primarykey"`
	path      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Blogs     []Blog `gorm:"many2many:blog_images"`
}
