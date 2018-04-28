package controller

import "github.com/labstack/echo"

type ArticleController struct {
	Base *BaseController
}

type ArticleForm struct {
	
}

func (ctx *ArticleController) AddArticle(c echo.Context) error {

	// service := service.CategoryService{Base: ctx.Base.Service}

	return BaseResponse(c, true, STATUS_OK, "add article successfully", nil)
}
