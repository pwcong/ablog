package model

type Evaluation struct {
	BaseModel
	IP        string  `json:"ip"`
	Score     int     `json:"score"`
	Content   string  `gorm:"type:text;" json:"content"`
	Article   Article `json:"article"`
	ArticleID uint    `json:"article_id"`
}
