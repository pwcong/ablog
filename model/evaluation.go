package model

import (
	"github.com/jinzhu/gorm"
)

type Evaluation struct {
	gorm.Model
	IP        string
	Score     int
	Content   string `gorm:"type:text;"`
	ArticleID uint
}
