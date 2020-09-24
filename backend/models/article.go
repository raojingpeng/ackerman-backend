package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title     string    `gorm:"column:title;size:255"`
	Desc      string    `gorm:"column:desc;size:255"`
	Content   string    `gorm:"column:content"`
	Timestamp time.Time `gorm:"column:timestamp;type:datetime"`
	Views     int       `gorm:"column:views"`
}
