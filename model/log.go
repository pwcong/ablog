package model

type Log struct {
	BaseModel
	Level  uint   `gorm:"not null" json:"level"`
	IP     string `gorm:"not null" json:"ip"`
	Action string `gorm:"not null" json:"action"`
}
