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

func (ctx *CategoryService) updateCategory(id uint, name string) (model.Category, error) {
	db := ctx.Base.DB

	var category model.Category

	notFound := db.Where("id = ?", id).First(&category).RecordNotFound()
	if notFound {
		return model.Category{}, errors.New("category is not existed")
	}

	category.Name = name

	return category, db.Update(category).Error

}

func (ctx *TagService) getCategory(id uint) (model.Category, error) {
	db := ctx.Base.DB

	var category model.Category

	notFound := db.Where("id = ?", id).First(&category).RecordNotFound()
	if notFound {
		return model.Category{}, errors.New("category is not existed")
	}

	return category, nil

}

func (ctx *CategoryService) getCategories() ([]model.Category, error) {
	db := ctx.Base.DB

	var categories []model.Category

	return categories, db.Find(&categories).Error

}

func (ctx *TagService) DelCategory(id uint) (model.Category, error) {

	db := ctx.Base.DB

	var category model.Category

	notFound := db.Where("id = ?", id).First(&category).RecordNotFound()
	if notFound {
		return model.Category{}, errors.New("category is not existed")
	}

	return category, db.Delete(&category).Error
}
