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

	category := model.Category{
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

	category.Name = name

	return category, db.Save(&category).Error

}

func (ctx *CategoryService) GetCategory(id uint) (model.Category, error) {
	db := ctx.Base.DB

	var category model.Category

	notFound := db.Where("id = ?", id).First(&category).RecordNotFound()
	if notFound {
		return model.Category{}, errors.New("category is not existed")
	}

	return category, nil

}

func (ctx *CategoryService) GetCategories(pageNo int, pageSize int) (model.Page, error) {

	db := ctx.Base.DB

	var totalSize int
	if err := db.Model(&model.Category{}).Count(&totalSize).Error; err != nil {
		return model.Page{}, err
	}

	var categories []model.Category
	if err := db.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&categories).Error; err != nil {
		return model.Page{}, err
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
