package service

import (
	"github.com/jinzhu/gorm"
	"github.com/pwcong/ablog/config"
	"github.com/pwcong/ablog/model"
)

type BaseService struct {
	Conf *config.Config
	DB   *gorm.DB
}

const (
	INFO = iota
	WARN
	ERROR
)

func (ctx *BaseService) Log(level uint, ip string, action string) {
	db := ctx.DB

	db.Create(&model.Log{
		Level:  level,
		IP:     ip,
		Action: action,
	})

}

func (ctx *BaseService) Info(ip string, action string) {
	ctx.Log(INFO, ip, action)
}
func (ctx *BaseService) Warn(ip string, action string) {
	ctx.Log(INFO, ip, action)
}
func (ctx *BaseService) Error(ip string, action string) {
	ctx.Log(INFO, ip, action)
}

func ConvertPageParameter(pageNo int, pageSize int) (int, int) {

	offset := (pageNo - 1) * pageSize

	if offset <= 0 {
		offset = -1
	}

	limit := pageSize
	if limit < 0 {
		limit = -1
	}

	return offset, limit
}
