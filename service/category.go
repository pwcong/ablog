package service

import (
	"errors"

	"github.com/pwcong/ablog/model"
)

type CategoryService struct {
	Base *BaseService
}

func (ctx *CategoryService) AddCategory(name string) (model.Category, error) {

	db := ctx.Base.DB

	var category model.Category

	notFound := db.Where("name = ?", name).First(&category).RecordNotFound()
	if !notFound {
		return category, nil
	}

	category = model.Category{
		Name: name,
	}

	return category, db.Create(&category).Error
}

func (ctx *CategoryService) UpdateCategory(id uint, name string) (model.Category, error) {
	db := ctx.Base.DB

	var category model.Category

	notFound := db.Where("id = ?", id).First(&category).RecordNotFound()
	if notFound {
		return model.Category{}, errors.New("category is not existed")
	}

	if category.Name == name {
		return category, nil
	}

	notFound = db.Where("name = ?", name).First(&model.Category{}).RecordNotFound()
	if !notFound {
		return model.Category{}, errors.New("category name is existed")
	}

	category.Name = name

	return category, db.Save(&category).Error

}

func (ctx *CategoryService) GetCategory(id uint) (model.Category, error) {
	db := ctx.Base.DB

	var category model.Category

	notFound := db.Preload("Articles").Where("id = ?", id).First(&category).RecordNotFound()
	if notFound {
		return model.Category{}, errors.New("category is not existed")
	}

	for i, _ := range category.Articles {
		category.Articles[i].Content = ""
	}

	return category, nil

}

func (ctx *CategoryService) GetCategories(pageNo int, pageSize int) (model.Page, error) {

	db := ctx.Base.DB

	var totalSize int
	if err := db.Table("categories").Where("deleted_at is null").Count(&totalSize).Error; err != nil {
		return model.Page{}, err
	}

	var categories []model.Category
	offset, limit := ConvertPageParameter(pageNo, pageSize)
	if err := db.Order("created_at desc").Preload("Articles").Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		return model.Page{}, err
	}

	for i, _ := range categories {
		for j, _ := range categories[i].Articles {
			categories[i].Articles[j].Content = ""
		}
	}

	return model.Page{
		PageNo:      pageNo,
		PageSize:    pageSize,
		CurrentSize: len(categories),
		TotalSize:   totalSize,
		Data:        categories,
	}, nil

}

func (ctx *CategoryService) DelCategory(id uint) error {

	db := ctx.Base.DB

	var category model.Category

	notFound := db.Where("id = ?", id).First(&category).RecordNotFound()
	if notFound {
		return errors.New("category is not existed")
	}

	tx := db.Begin()

	if err := tx.Model(&category).Association("Articles").Clear().Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&category).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}
