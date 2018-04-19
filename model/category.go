package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);unique_index;not null;"`
	Articles []Article
}
