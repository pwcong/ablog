package service

import (
	"github.com/pwcong/ablog/model"
)

type CounterService struct {
	Base *BaseService
}

func (ctx *CounterService) PlusCounter(key string, id uint) model.Counter {

	db := ctx.Base.DB

	var counter model.Counter

	counter = ctx.GetCounter(key, id)

	counter.Count++

	db.Save(&counter)

	return counter

}

func (ctx *CounterService) GetCounter(key string, id uint) model.Counter {

	db := ctx.Base.DB

	var counter model.Counter

	db.Where(model.Counter{
		Key:      key,
		TargetID: id,
	}).FirstOrCreate(&counter)

	return counter

}
