package model

type Article struct {
	BaseModel
	Title       string       `json:"title"`
	Content     string       `gorm:"type:longtext;PRELOAD:false" json:"content"`
	Banner      string       `json:"banner"`
	Category    Category     `json:"category"`
	CategoryID  uint         `json:"category_id"`
	Tags        []Tag        `gorm:"many2many:article_tags;" json:"tags"`
	Evaluations []Evaluation `json:"evaluations"`
}
