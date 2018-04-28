package service

import (
	"errors"

	"github.com/pwcong/ablog/model"
)

type ArticleService struct {
	Base *BaseService
}

func (ctx *ArticleService) DelArticle(id uint) error {

	db := ctx.Base.DB

	var article model.Article

	notFound := db.Where("id = ?", id).First(&article).RecordNotFound()
	if notFound {
		return errors.New("article is not existed")
	}

	tx := db.Begin()

	if err := tx.Model(&article).Association("Tags").Clear().Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&article).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
	if notFound {
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

func (ctx *ArticleService) UpdateArticle(id uint, title string, content string, banner string, categoryId uint, tagIds []uint) (model.Article, error) {

	db := ctx.Base.DB

	var article model.Article
	notFound := db.Where("id = ?", id).First(&article).RecordNotFound()
	if notFound {
		return model.Article{}, errors.New("article is not existed")
	}

	tx := db.Begin()
	article.Title = title
	article.Content = content
	article.Banner = banner
	if err := tx.Save(&article).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	var category model.Category
	notFound = db.Where("id = ?", categoryId).First(&category).RecordNotFound()
	if notFound {
		return model.Article{}, errors.New("category is not existed")
	}

	if err := tx.Model(&article).Association("Category").Replace(category).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	var tags []model.Tag
	db.Where("id in (?)", tagIds).Find(&tags)
	if err := tx.Model(&article).Association("Tags").Replace(tags).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	return article, tx.Commit().Error
}

func (ctx *ArticleService) GetArticle(id uint) (model.Article, error) {

	db := ctx.Base.DB

	var article model.Article
	notFound := db.Where("id = ?", id).First(&article).RecordNotFound()
	if notFound {
		return model.Article{}, errors.New("article is not existed")
	}

	if err := db.Preload("Category").First(&article).Error; err != nil {
		return model.Article{}, err
	}

	if err := db.Preload("Tags").First(&article).Error; err != nil {
		return model.Article{}, err
	}

	return article, nil

}

func (ctx *ArticleService) GetArticlesByCategoryId(categoryId uint, pageNo int, pageSize int) (model.Page, error) {

	db := ctx.Base.DB

	var category model.Category
	if notFound := db.Where("id = ?", categoryId).First(&category).RecordNotFound(); notFound {
		return model.Page{}, errors.New("category is not existed")
	}

	totalSize := db.Model(&category).Association("Articles").Count()

	var articles []model.Article
	if err := db.Model(&category).Offset((pageNo-1)*pageSize).Limit(pageSize).Related(&articles, "Articles").Error; err != nil {
		return model.Page{}, err
	}

	return model.Page{
		PageNo:      pageNo,
		PageSize:    pageSize,
		CurrentSize: len(articles),
		TotalSize:   totalSize,
		Data:        articles,
	}, nil

}

func (ctx *ArticleService) GetArticlesByTagId(tagId uint, pageNo int, pageSize int) (model.Page, error) {

	db := ctx.Base.DB

	var tag model.Tag
	if notFound := db.Where("id = ?", tagId).First(&tag).RecordNotFound(); notFound {
		return model.Page{}, errors.New("tag is not existed")
	}

	totalSize := db.Model(&tag).Association("Articles").Count()

	var articles []model.Article
	if err := db.Model(&tag).Offset((pageNo-1)*pageSize).Limit(pageSize).Related(&articles, "Articles").Error; err != nil {
		return model.Page{}, err
	}

	return model.Page{
		PageNo:      pageNo,
		PageSize:    pageSize,
		CurrentSize: len(articles),
		TotalSize:   totalSize,
		Data:        articles,
	}, nil

}
