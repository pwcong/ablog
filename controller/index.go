package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type IndexController struct {
	Base *BaseController
}

func (ctx *IndexController) Default(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Hello": "World!",
	})
}
