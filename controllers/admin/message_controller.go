package admin

import (
	"github.com/kataras/iris"
	"github.com/mlogclub/mlog/model"
	"github.com/mlogclub/mlog/services"
	"github.com/mlogclub/simple"
	"strconv"
)

type MessageController struct {
	Ctx            iris.Context
	MessageService *services.MessageService
}

func (this *MessageController) GetBy(id int64) *simple.JsonResult {
	t := this.MessageService.Get(id)
	if t == nil {
		return simple.ErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *MessageController) AnyList() *simple.JsonResult {
	list, paging := this.MessageService.Query(simple.NewParamQueries(this.Ctx).PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *MessageController) PostCreate() *simple.JsonResult {
	t := &model.Message{}
	this.Ctx.ReadForm(t)

	err := this.MessageService.Create(t)
	if err != nil {
		return simple.ErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *MessageController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt64(this.Ctx, "id")
	if err != nil {
		return simple.ErrorMsg(err.Error())
	}
	t := this.MessageService.Get(id)
	if t == nil {
		return simple.ErrorMsg("entity not found")
	}

	this.Ctx.ReadForm(t)

	err = this.MessageService.Update(t)
	if err != nil {
		return simple.ErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}
