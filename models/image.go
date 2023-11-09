package models

import "time"

type Image struct {
	ID        uint   `gorm:"primarykey"`
	Path      string `json:"path" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Blogs     []Blog `gorm:"many2many:blog_images;" json:"blogs"`
}
