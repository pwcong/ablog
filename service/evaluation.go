package service

import (
	"errors"

	"github.com/pwcong/ablog/model"
)

type EvaluationService struct {
	Base *BaseService
}

func (ctx *EvaluationService) AddEvaluation(ip string, score int, content string, articleID uint) (model.Evaluation, error) {

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

func (ctx *EvaluationService) GetEvaluations(articleID uint) ([]model.Evaluation, error) {
	db := ctx.Base.DB

	var evaluations []model.Evaluation

	if err := db.Where("article_id = ?", articleID).Find(&evaluations).Error; err != nil {
		return []model.Evaluation{}, err
	}

	return evaluations, nil

}
