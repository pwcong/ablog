package model

type Tag struct {
	BaseModel
	Name     string    `gorm:"type:varchar(255);unique_index;not null;" json:"name"`
	Articles []Article `gorm:"many2many:article_tags;" json:"articles"`
}
