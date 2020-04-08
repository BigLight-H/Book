package controllers

import (
	"fiction_web/util"
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
	if p.controllerName == "user"{
		token := strings.Fields(p.Ctx.Input.Header("Authorization"))
		err := util.ValidateToken(token[1])
		if err != nil {
			p.MsgBack("token不存在,请重新登录!", -1)
			panic("token不存在,请重新登录!")
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