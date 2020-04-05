package controllers

import (
	"fiction_web/models"
	"fiction_web/service"
	"fiction_web/util"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//搜索首页
func (p *MainController) Home() {
	s := p.GetString(":name")
	info := service.GetBook(s)
	p.Data["json"] = util.ReturnJson("首页数据", 200, info)
	p.ServeJSON()
}

//书目录页
func (p *MainController) List() {
	links := p.GetString("link")
	id := util.StrToInt(p.GetString("id"))
	list := models.Lists{Id:id}
	p.o.Read(&list)
	info := service.BookList(list, id, links)
	p.Data["json"] = util.ReturnJson("目录", 200, info)
	p.ServeJSON()
}


//内容页
func (p *MainController) Detail() {
	links := p.GetString("link")
	id := util.StrToInt(p.GetString("id"))
	con := models.Content{Id:id}
	p.o.Read(&con)
	info := service.BookContent(con, links)
	p.Data["json"] = util.ReturnJson("详情内容页", 200, info)
	p.ServeJSON()
}

//书简介
func (p *MainController) Synopsis() {
	links := p.GetString("link")
	id := util.StrToInt(p.GetString("id"))
	con := models.Synopsis{Id:id}
	p.o.Read(&con)
	info := service.BookSynosis(con, links)
	p.Data["json"] = util.ReturnJson("文章目录介绍部分", 200, info)
	p.ServeJSON()
}


