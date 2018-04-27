package model

type Article struct {
	BaseModel
	Title       string
	Banner      string
	Category    Category
	CategoryID  uint
	Tags        []Tag  `gorm:"many2many:article_tags;"`
	Content     string `gorm:"type:longtext;"`
	Evaluations []Evaluation
}
