package model

type Evaluation struct {
	BaseModel
	IP        string
	Score     int
	Content   string `gorm:"type:text;"`
	ArticleID uint
}
