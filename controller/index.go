package controller

import (
	"github.com/labstack/echo"
)

type IndexController struct {
	Base *BaseController
}

func (ctx *IndexController) Default(c echo.Context) error {
	return c.File("view/index.html")
}
