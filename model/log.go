package model

import "github.com/jinzhu/gorm"

type Log struct {
	gorm.Model
	Level  uint   `gorm:"not null"`
	IP     string `gorm:"not null"`
	Action string `gorm:"not null"`
}
