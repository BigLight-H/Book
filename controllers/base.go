package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseController struct {
	beego.Controller
	o orm.Ormer
}

//构造函数
func (p *BaseController) Prepare()  {
	p.o = orm.NewOrm()
}