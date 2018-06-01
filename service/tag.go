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

	var tag model.Tag

	notFound := db.Where("name = ?", name).First(&tag).RecordNotFound()
	if !notFound {
		return tag, nil
	}

	tag = model.Tag{
		Name: name,
	}

	return tag, db.Create(&tag).Error
}

func (ctx *TagService) UpdateTag(id uint, name string) (model.Tag, error) {
	db := ctx.Base.DB

	var tag model.Tag

	notFound := db.Where("id = ?", id).First(&tag).RecordNotFound()
	if notFound {
		return model.Tag{}, errors.New("tag is not existed")
	}

	if tag.Name == name {
		return tag, nil
	}

	notFound = db.Where("name = ?", name).First(&model.Tag{}).RecordNotFound()
	if !notFound {
		return model.Tag{}, errors.New("tag name is existed")
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
	if err := db.Table("tags").Where("deleted_at is not null").Count(&totalSize).Error; err != nil {
		return model.Page{}, err
	}

	var tags []model.Tag
	offset, limit := ConvertPageParameter(pageNo, pageSize)
	if err := db.Offset(offset).Limit(limit).Find(&tags).Error; err != nil {
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
