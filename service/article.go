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

	if err := tx.Model(&article).Association("Category").Clear().Error; err != nil {
		tx.Rollback()
		return err
	}

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

func (ctx *ArticleService) AddArticle(title string, content string, banner string, categoryID uint, tagIDs []uint) (model.Article, error) {

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
	notFound := db.Where("id = ?", categoryID).First(&category).RecordNotFound()
	if notFound {
		return model.Article{}, errors.New("category is not existed")
	}

	if err := tx.Model(&article).Association("Category").Append(category).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	var tags []model.Tag
	db.Where("id in (?)", tagIDs).Find(&tags)
	if err := tx.Model(&article).Association("Tags").Append(tags).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	return article, tx.Commit().Error
}

func (ctx *ArticleService) UpdateArticle(id uint, title string, content string, banner string, categoryID uint, tagIDs []uint) (model.Article, error) {

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
	notFound = db.Where("id = ?", categoryID).First(&category).RecordNotFound()
	if notFound {
		return model.Article{}, errors.New("category is not existed")
	}

	if err := tx.Model(&article).Association("Category").Replace(category).Error; err != nil {
		tx.Rollback()
		return model.Article{}, err
	}

	var tags []model.Tag
	db.Where("id in (?)", tagIDs).Find(&tags)
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

	if err := db.Preload("Category").Preload("Tags").First(&article).Error; err != nil {
		return model.Article{}, err
	}

	return article, nil

}

func (ctx *ArticleService) GetArticlesByCategoryID(categoryID uint, pageNo int, pageSize int) (model.Page, error) {

	db := ctx.Base.DB

	var category model.Category
	if notFound := db.Where("id = ?", categoryID).First(&category).RecordNotFound(); notFound {
		return model.Page{}, errors.New("category is not existed")
	}

	totalSize := db.Model(&category).Association("Articles").Count()

	offset, limit := ConvertPageParameter(pageNo, pageSize)

	var articles []model.Article
	if err := db.Model(&category).Offset(offset).Limit(limit).Related(&articles, "Articles").Error; err != nil {
		return model.Page{}, err
	}

	if err := db.Preload("Category").Preload("Tags").Find(&articles).Error; err != nil {
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

func (ctx *ArticleService) GetArticlesByTagID(tagID uint, pageNo int, pageSize int) (model.Page, error) {

	db := ctx.Base.DB

	var tag model.Tag
	if notFound := db.Where("id = ?", tagID).First(&tag).RecordNotFound(); notFound {
		return model.Page{}, errors.New("tag is not existed")
	}

	totalSize := db.Model(&tag).Association("Articles").Count()

	offset, limit := ConvertPageParameter(pageNo, pageSize)

	var articles []model.Article
	if err := db.Model(&tag).Offset(offset).Limit(limit).Related(&articles, "Articles").Error; err != nil {
		return model.Page{}, err
	}

	if err := db.Preload("Category").Preload("Tags").Find(&articles).Error; err != nil {
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

func (ctx *ArticleService) GetArticles(pageNo int, pageSize int) (model.Page, error) {

	db := ctx.Base.DB

	var totalSize int
	if err := db.Table("articles").Where("deleted_at is not null").Count(&totalSize).Error; err != nil {
		return model.Page{}, err
	}

	var articles []model.Article

	offset, limit := ConvertPageParameter(pageNo, pageSize)

	if err := db.Offset(offset).Limit(limit).Find(&articles).Error; err != nil {
		return model.Page{}, err
	}

	if err := db.Preload("Category").Preload("Tags").Find(&articles).Error; err != nil {
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
