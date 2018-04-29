package controller

import (
	"errors"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pwcong/ablog/service"
)

type EvaluationForm struct {
	PageSize int    `query:"page_size"`
	PageNo   int    `query:"page_no"`
	Score    int    `json:"score" form:"score"`
	Content  string `json:"content" form:"content"`
}

type EvaluationController struct {
	Base *BaseController
}

func (ctx *EvaluationController) GetEvaluation(c echo.Context) error {

	service := service.EvaluationService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	evaluation, err := service.GetEvaluation(uint(id))

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get evaluation successfully", evaluation)

}

func (ctx *EvaluationController) GetEvaluations(c echo.Context) error {

	service := service.EvaluationService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	form := new(EvaluationForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	pageNo, pageSize := ResolvePageParameter(form.PageNo, form.PageSize)

	page, err := service.GetEvaluations(uint(id), pageNo, pageSize)

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get evaluations successfully", page)

}

func (ctx *EvaluationController) AddEvaluation(c echo.Context) error {

	service := service.EvaluationService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	form := new(EvaluationForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	if form.Content == "" {
		return errors.New("content can not be empty")
	}

	evaluation, err := service.AddEvaluation(uint(id), c.RealIP(), form.Score, form.Content)

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "add evaluation successfully", evaluation)

}

func (ctx *EvaluationController) DelEvaluation(c echo.Context) error {

	service := service.EvaluationService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	err = service.DelEvaluation(uint(id))

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "delete evaluation successfully", nil)

}
