package controllers

import (
	"fiction_web/models"
	"fiction_web/service"
	"github.com/davecgh/go-spew/spew"
)

type BookSourceController struct {
	BaseController
}

//获取首页数据
func (p *BookSourceController) BookSource() {
	source := models.Source{Id:1}
	p.o.Read(&source)
	lists := service.HubSource(source)
	spew.Dump(lists)
	p.Ctx.WriteString("hahha")
}

//获取分类页分类标题数据
func (p *BookSourceController) BookTypes() {
	types := []models.Types{}
	p.o.QueryTable(new(models.Types).TableName()).All(&types)
	p.Data["json"] = types
	p.ServeJSON()
}

//获取书源分类
func (p *BookSourceController) BookType() {
	link := p.GetString("link")
	t := models.Type{Id:1}
	p.o.Read(&t)
	data := service.BookType(t, link)
	p.Data["json"] = data
	p.ServeJSON()
}

//获取排行链接,名
func (p *BookSourceController) Leader() {
	leader := []models.Leader{}
	p.o.QueryTable(new(models.Leader).TableName()).All(&leader)
	p.Data["json"] = leader
	p.ServeJSON()
}

//获取指定链接排行榜数据
func (p *BookSourceController) LeaderBoard() {
	link := p.GetString("link")
	t := models.Leaderboard{Id:1}
	p.o.Read(&t)
	data := service.GetBoard(t, link)
	p.Data["json"] = data
	p.ServeJSON()
}

//获取指定链接的完本数据
func (p *BookSourceController) BookEnd() {
	link := "https://m.kuxiaoshuo.com/full/1/"
	if p.GetString("link") != "" {
		link = p.GetString("link")
	}
	t := models.Completed{Id:1}
	p.o.Read(&t)
	data := service.GetBookEnd(t, link)
	p.Data["json"] = data
	p.ServeJSON()
}









