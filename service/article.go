package service

import (
	"errors"

	"github.com/pwcong/ablog/model"
)

type ArticleService struct {
	Base *BaseService
}

func (ctx *ArticleService) AddArticle(title string, content string, banner string, categoryId uint, tagIds []uint) (model.Article, error) {

	db := ctx.Base.DB

	article := model.Article{
		Title:   title,
		Content: content,
		Banner:  banner,
	}

	tx := db.Begin()
	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	var category model.Category
	notFound := db.Where("id = ?", categoryId).First(&category).RecordNotFound()
	if !notFound {
		return model.Article{}, errors.New("category is not existed")
	}

	if err := tx.Model(&article).Association("Category").Append(category).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	var tags []model.Tag
	db.Where("id in (?)", tagIds).Find(&tags)
	if err := tx.Model(&article).Association("Tags").Append(tags).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	return article, tx.Commit().Error
}
