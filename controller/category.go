package controller

import (
	"errors"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pwcong/ablog/service"
)

type CategoryForm struct {
	PageSize int    `query:"page_size"`
	PageNo   int    `query:"page_no"`
	Name     string `json:"name" form:"name"`
}

type CategoryController struct {
	Base *BaseController
}

func (ctx *CategoryController) GetCategories(c echo.Context) error {

	service := service.CategoryService{Base: ctx.Base.Service}

	form := new(CategoryForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	pageNo, pageSize := ResolvePageParameter(form.PageNo, form.PageSize)

	page, err := service.GetCategories(pageNo, pageSize)
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get categories successfully", page)
}

func (ctx *CategoryController) GetCategory(c echo.Context) error {

	service := service.CategoryService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	category, err := service.GetCategory(uint(id))
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get category successfully", category)
}

func (ctx *CategoryController) DelCategory(c echo.Context) error {

	service := service.CategoryService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	err = service.DelCategory(uint(id))
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "delete category successfully", nil)
}

func (ctx *CategoryController) UpdateCategory(c echo.Context) error {

	service := service.CategoryService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	form := new(CategoryForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	if form.Name == "" {
		return errors.New("lack of name")
	}

	category, err := service.UpdateCategory(uint(id), form.Name)
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "update category successfully", category)
}

func (ctx *CategoryController) AddCategory(c echo.Context) error {

	service := service.CategoryService{Base: ctx.Base.Service}

	form := new(CategoryForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	if form.Name == "" {
		return errors.New("lack of name")
	}

	category, err := service.AddCategory(form.Name)
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "add category successfully", category)
}
