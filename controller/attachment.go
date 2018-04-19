package controller

import (
	"github.com/labstack/echo"
	"github.com/pwcong/ablog/model"
	"github.com/pwcong/ablog/service"
)

type AttachmentController struct {
	Base *BaseController
}

func (ctx *AttachmentController) Upload(c echo.Context) error {

	service := service.AttachmentService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "upload image")

	file, err := c.FormFile("file")
	if file == nil || err != nil {
		return BaseResponse(c, false, STATUS_ERROR, err.Error(), struct{}{})
	}

	var att model.Attachment

	att, err = service.SaveAttachment(file)

	if err != nil {
		return BaseResponse(c, false, STATUS_ERROR, err.Error(), struct{}{})
	}

	return BaseResponse(c, true, STATUS_OK, "upload successfully", struct {
		ID  uint   `json:"id"`
		URL string `json:"url"`
	}{
		ID:  att.ID,
		URL: "/public/" + att.Year + "/" + att.Month + "/" + att.Date + "/" + att.Symbol + "." + att.ExtName,
	})

}
