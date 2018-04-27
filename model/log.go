package model

type Log struct {
	BaseModel
	Level  uint   `gorm:"not null"`
	IP     string `gorm:"not null"`
	Action string `gorm:"not null"`
}
