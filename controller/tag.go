package controller

import (
	"errors"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pwcong/ablog/service"
)

type TagForm struct {
	PageSize int    `query:"page_size"`
	PageNo   int    `query:"page_no"`
	Name     string `json:"name" form:"name"`
}

type TagController struct {
	Base *BaseController
}

func (ctx *TagController) GetTags(c echo.Context) error {

	service := service.TagService{Base: ctx.Base.Service}

	form := new(TagForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	pageNo, pageSize := ResolvePageParameter(form.PageNo, form.PageSize)

	page, err := service.GetTags(pageNo, pageSize)
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get tags successfully", page)
}

func (ctx *TagController) GetTag(c echo.Context) error {

	service := service.TagService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	tag, err := service.GetTag(uint(id))
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get tag successfully", tag)
}

func (ctx *TagController) DelTag(c echo.Context) error {

	service := service.TagService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	err = service.DelTag(uint(id))
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "delete tag successfully", nil)
}

func (ctx *TagController) UpdateTag(c echo.Context) error {

	service := service.TagService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	form := new(TagForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	if form.Name == "" {
		return errors.New("lack of name")
	}

	tag, err := service.UpdateTag(uint(id), form.Name)
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "update tag successfully", tag)
}

func (ctx *TagController) AddTag(c echo.Context) error {

	service := service.TagService{Base: ctx.Base.Service}

	form := new(TagForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	if form.Name == "" {
		return errors.New("lack of name")
	}

	tag, err := service.AddTag(form.Name)
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "add tag successfully", tag)
}
