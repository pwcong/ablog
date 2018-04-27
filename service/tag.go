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

func (ctx *TagService) UpdateTag(id uint, name string) (model.Tag, error) {
	db := ctx.Base.DB

	var tag model.Tag

	notFound := db.Where("id = ?", id).First(&tag).RecordNotFound()
	if notFound {
		return model.Tag{}, errors.New("tag is not existed")
	}

	tag.Name = name

	return tag, db.Save(tag).Error

}

func (ctx *TagService) GetTag(id uint) (model.Tag, error) {
	db := ctx.Base.DB

	var tag model.Tag

	notFound := db.Where("id = ?", id).First(&tag).RecordNotFound()
	if notFound {
		return model.Tag{}, errors.New("tag is not existed")
	}

	return tag, nil

}

func (ctx *TagService) GetTags(pageNo int, pageSize int) (model.Page, error) {
	db := ctx.Base.DB

	var totalSize int
	if err := db.Model(&model.Category{}).Count(&totalSize).Error; err != nil {
		return model.Page{}, err
	}

	var tags []model.Tag
	if err := db.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&tags).Error; err != nil {
		return model.Page{}, err
	}

	return model.Page{
		PageNo:      pageNo,
		PageSize:    pageSize,
		CurrentSize: len(tags),
		TotalSize:   totalSize,
		Data:        tags,
	}, nil

}

func (ctx *TagService) DelTag(id uint) error {

	db := ctx.Base.DB

	var tag model.Tag

	notFound := db.Where("id = ?", id).First(&tag).RecordNotFound()
	if notFound {
		return errors.New("tag is not existed")
	}

	tx := db.Begin()

	if err := tx.Model(&tag).Association("Articles").Clear().Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&tag).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
