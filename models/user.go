package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
	blog      []Blog `gorm:"many2many:blog_users"`
}
