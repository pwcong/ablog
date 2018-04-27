package controller

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/pwcong/ablog/service"
)

type CategoryForm struct {
	PageSize int `json:"page_size" xml:"page_size" form:"page_size" query:"page_size"`
	PageNo   int `json:"page_no" xml:"page_no" form:"page_no" query:"page_no"`
}

type CategoryController struct {
	Base *BaseController
}

func (ctx *CategoryController) GetCategories(c echo.Context) error {

	service := service.CategoryService{Base: ctx.Base.Service}

	form := new(CategoryForm)
	if err := c.Bind(form); err != nil {
		return errors.New("invalid params")
	}

	pageNo, pageSize := ResolvePageParameter(form.PageNo, form.PageSize)

	page, err := service.GetCategories(pageNo, pageSize)
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get categories successfully", page)
}
