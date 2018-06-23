package controller

import (
	"github.com/labstack/echo"
	"github.com/pwcong/ablog/model"
	"github.com/pwcong/ablog/service"
)

type AttachmentController struct {
	Base *BaseController
}

type AttachmentForm struct {
	Symbol string `form:"symbol"`
}

func (ctx *AttachmentController) Upload(c echo.Context) error {

	service := service.AttachmentService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "upload image")

	file, err := c.FormFile("file")
	if file == nil || err != nil {
		return err
	}

	form := new(AttachmentForm)
	if err = c.Bind(form); err != nil {
		return err
	}

	var att model.Attachment
	if form.Symbol != "" {
		_att, _err := service.GetAttachmentBySymbol(form.Symbol)
		if _err != nil {
			att, err = service.SaveAttachment(file)
		} else {
			att = _att
		}

	} else {
		att, err = service.SaveAttachment(file)
	}

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "upload successfully", struct {
		ID  uint   `json:"id"`
		URL string `json:"url"`
	}{
		ID:  att.ID,
		URL: "/public/" + att.Year + "/" + att.Month + "/" + att.Date + "/" + att.Symbol + "." + att.ExtName,
	})

}
