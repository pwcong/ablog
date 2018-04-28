package model

type Attachment struct {
	BaseModel
	Symbol   string `gorm:"not null" json:"symbol"`
	Filename string `gorm:"not null" json:"filename"`
	Year     string `gorm:"not null" json:"year"`
	Month    string `gorm:"not null" json:"month"`
	Date     string `gorm:"not null" json:"date"`
	ExtName  string `gorm:"not null" json:"ext_name"`
}
