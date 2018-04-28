package controller

import (
	"errors"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pwcong/ablog/model"
	"github.com/pwcong/ablog/service"
)

type ArticleController struct {
	Base *BaseController
}

type ArticleForm struct {
	PageSize   int    `query:"page_size"`
	PageNo     int    `query:"page_no"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	Banner     string `json:"banner" form:"banner"`
	CategoryID uint   `json:"category_id" form:"category_id"`
	TagIDs     []uint `json:"tag_ids" form:"tag_ids"`
}

func (ctx *ArticleController) AddArticle(c echo.Context) error {

	service := service.ArticleService{Base: ctx.Base.Service}

	form := new(ArticleForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	article, err := service.AddArticle(form.Title, form.Content, form.Banner, form.CategoryID, form.TagIDs)

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "add article successfully", article)
}

func (ctx *ArticleController) GetArticle(c echo.Context) error {

	service := service.ArticleService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	article, err := service.GetArticle(uint(id))

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get article successfully", article)
}

func (ctx *ArticleController) GetArticlesByFlagWithId(c echo.Context) error {

	service := service.ArticleService{Base: ctx.Base.Service}

	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	form := new(ArticleForm)
	if err := c.Bind(form); err != nil {
		return err
	}
	pageNo, pageSize := ResolvePageParameter(form.PageNo, form.PageSize)

	flag := c.Param("flag")
	if flag != "category" && flag != "tag" {
		return errors.New("invalid flag")
	}

	var page model.Page
	if flag == "category" {
		page, err = service.GetArticlesByCategoryId(uint(id), pageNo, pageSize)
	} else if flag == "tag" {
		page, err = service.GetArticlesByTagId(uint(id), pageNo, pageSize)
	}

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get articles successfully", page)
}
