package service

import (
	"errors"

	"github.com/pwcong/ablog/model"
)

type EvaluationService struct {
	Base *BaseService
}

func (ctx *EvaluationService) AddEvaluation(articleID uint, ip string, score int, content string) (model.Evaluation, error) {

	db := ctx.Base.DB

	var article model.Article
	if notFound := db.Where("id = ?", articleID).First(&article).RecordNotFound(); notFound {
		return model.Evaluation{}, errors.New("article is not existed")
	}

	evaluation := model.Evaluation{
		IP:      ip,
		Score:   score,
		Content: content,
	}

	tx := db.Begin()

	if err := tx.Create(&evaluation).Error; err != nil {
		tx.Rollback()
		return model.Evaluation{}, err
	}

	if err := tx.Model(&evaluation).Association("Article").Append(article).Error; err != nil {
		tx.Rollback()
		return model.Evaluation{}, err
	}

	return evaluation, tx.Commit().Error
}

func (ctx *EvaluationService) GetEvaluation(id uint) (model.Evaluation, error) {
	db := ctx.Base.DB

	var evaluation model.Evaluation

	notFound := db.Where("id = ?", id).First(&evaluation).RecordNotFound()
	if notFound {
		return model.Evaluation{}, errors.New("evaluation is not existed")
	}

	return evaluation, nil

}

func (ctx *EvaluationService) GetEvaluations(articleID uint, pageNo int, pageSize int) (model.Page, error) {
	db := ctx.Base.DB

	var article model.Article
	if notFound := db.Where("id = ?", articleID).First(&article).RecordNotFound(); notFound {
		return model.Page{}, errors.New("article is not existed")
	}

	totalSize := db.Model(&article).Association("Evaluations").Count()
	var evaluations []model.Evaluation
	offset, limit := ConvertPageParameter(pageNo, pageSize)
	if err := db.Model(&article).Offset(offset).Limit(limit).Related(&evaluations, "Evaluations").Error; err != nil {
		return model.Page{}, err
	}

	return model.Page{
		PageNo:      pageNo,
		PageSize:    pageSize,
		CurrentSize: len(evaluations),
		TotalSize:   totalSize,
		Data:        evaluations,
	}, nil

}

func (ctx *EvaluationService) DelEvaluation(id uint) error {
	db := ctx.Base.DB

	var evaluation model.Evaluation
	if notFound := db.Where("id = ?", id).First(&evaluation).RecordNotFound(); notFound {
		return errors.New("evaluation is not existed")
	}

	tx := db.Begin()
	if err := tx.Model(&evaluation).Association("Article").Clear().Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&evaluation).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}
