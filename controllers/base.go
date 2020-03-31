package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseController struct {
	beego.Controller
	o orm.Ormer
}

//Json结构体
type Json struct {
	Msg string
	Status int
}

//构造函数
func (p *BaseController) Prepare()  {
	p.o = orm.NewOrm()
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