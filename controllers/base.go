package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type BaseController struct {
	beego.Controller
	o orm.Ormer
	controllerName string
	actionName     string
}

//Json结构体
type Json struct {
	Msg string
	Status int
}

//构造函数
func (p *BaseController) Prepare()  {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
	if p.controllerName == "users"{
		if p.GetSession("user") == nil {
			p.MsgBack("未登录", 0)
		} else {
			p.Data["user_data"] = p.GetSession("user")
			p.Data["user_id"] = p.GetSession("user_id")
		}
	}
}

//返回json信息
func (p *BaseController) MsgBack(msg string, status int)  {
	data := &Json{
		msg,
		status,
	}
	p.Data["json"] = data
	p.ServeJSON()
}