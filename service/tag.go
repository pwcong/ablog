package service

import (
	"errors"

	"github.com/pwcong/ablog/model"
)

type TagService struct {
	Base *BaseService
}

func (ctx *TagService) AddTag(name string) (model.Tag, error) {

	db := ctx.Base.DB

	tag := model.Tag{
		Name: name,
	}

	if dbc := db.Create(&tag); dbc.Error != nil {
		return model.Tag{}, dbc.Error
	}

	return tag, nil
}

func (ctx *TagService) updateTag(id uint, name string) (model.Tag, error) {
	db := ctx.Base.DB

	var tag model.Tag

	notFound := db.Where("id = ?", id).First(&tag).RecordNotFound()
	if notFound {
		return model.Tag{}, errors.New("tag is not existed")
	}

	tag.Name = name

	if err := db.Update(tag).Error; err != nil {
		return model.Tag{}, err
	}

	return tag, nil

}

func (ctx *TagService) getTag(id uint) (model.Tag, error) {
	db := ctx.Base.DB

	var tag model.Tag

	notFound := db.Where("id = ?", id).First(&tag).RecordNotFound()
	if notFound {
		return model.Tag{}, errors.New("tag is not existed")
	}

	return tag, nil

}

func (ctx *TagService) getTags() ([]model.Tag, error) {
	db := ctx.Base.DB

	var tags []model.Tag

	if err := db.Find(&tags).Error; err != nil {
		return []model.Tag{}, err
	}

	return tags, nil

}

func (ctx *TagService) DelTag(id uint) (model.Tag, error) {

	db := ctx.Base.DB

	var tag model.Tag

	notFound := db.Where("id = ?", id).First(&tag).RecordNotFound()
	if notFound {
		return model.Tag{}, errors.New("tag is not existed")
	}

	tx := db.Begin()
	if err := tx.Delete(&tag).Error; err != nil {
		tx.Rollback()
		return model.Tag{}, err
	}

	if err := tx.Model(&tag).Association("Articles").Delete(tag).Error; err != nil {
		tx.Rollback()
		return model.Tag{}, err
	}

	return tag, tx.Commit().Error
}