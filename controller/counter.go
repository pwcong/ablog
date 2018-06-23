package controller

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/pwcong/ablog/service"
)

type CounterController struct {
	Base *BaseController
}

type CounterForm struct {
	Key string `json:"key" form:"key"`
	ID  uint   `json:"id" form:"id"`
}

func (ctx *CounterController) PlusCounter(c echo.Context) error {

	service := service.CounterService{Base: ctx.Base.Service}

	form := new(CounterForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	counter := service.PlusCounter(form.Key, form.ID)

	return BaseResponse(c, true, STATUS_OK, "plus counter successfully", counter)
}

func (ctx *CounterController) GetCounter(c echo.Context) error {

	service := service.CounterService{Base: ctx.Base.Service}

	key := c.Param("key")
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	counter := service.GetCounter(key, uint(id))

	return BaseResponse(c, true, STATUS_OK, "get counter successfully", counter)

}
