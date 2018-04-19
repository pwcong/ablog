package model

import (
	"github.com/jinzhu/gorm"
)

type Attachment struct {
	gorm.Model
	Symbol   string `gorm:"not null;unique"`
	Filename string `gorm:"not null"`
	Year     string `gorm:"not null"`
	Month    string `gorm:"not null"`
	Date     string `gorm:"not null"`
	ExtName  string `gorm:"not null"`
}
