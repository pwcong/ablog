package model

type Category struct {
	BaseModel
	Name     string    `gorm:"type:varchar(255);unique_index;not null;" json:"name"`
	Articles []Article `json:"articles"`
}
