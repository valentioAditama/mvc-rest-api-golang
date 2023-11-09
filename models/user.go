package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Blogs     []*Blog `gorm:"many2many:blog_users;" json:"blogs"`
}
