package users

import (
	"fiction_web/controllers"
	"fiction_web/models"
	"fiction_web/util"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	controllers.BaseController
	o orm.Ormer
}

//添加书架
func (u *UserController) AddBookshelf() {
	hub_id := u.GetString("hub_id")
	link := u.GetString("link")
	domain := u.GetString("domain")
	img := u.GetString("img")
	author := u.GetString("author")
	book_name := u.GetString("book_name")
	user_id := u.Data["user_id"].(int)
	book := models.Bookshelf{}
	book.Link = link
	book.UserId = user_id
	book.HubId = util.StrToInt(hub_id)
	book.Domain = domain
	book.BookName = book_name
	book.Img = img
	book.Author = author
	u.o.Insert(&book)
}

//查看书架
func (u *UserController) GetBookshelf() {
	user_id := u.Data["user_id"].(int)
	book := []models.Bookshelf{}
	u.o.QueryTable(new(models.Bookshelf).TableName()).Filter("user_id", user_id).All(&book)
	u.Data["json"] = book
	u.ServeJSON()
}

//查看该书本是否已经加入书架
func (u *UserController) VerificationBook() {
	user_id := u.Data["user_id"].(int)
	domain := u.GetString("domain")
	num,err := u.o.QueryTable(new(models.Bookshelf).TableName()).
		Filter("user_id", user_id).Filter("domain", domain).
		Count()
	if err ==nil {
		if num == 0 {
			u.MsgBack("已存在", 1)
		}
	}
	u.MsgBack("不存在", 0)
}

//更新图书阅读进度
func (u *UserController) UpdateBookSchedule() {
	user_id := u.Data["user_id"].(int)
	domain := u.GetString("domain")
	link := u.GetString("link")
	_, err := u.o.QueryTable(new(models.Bookshelf).TableName()).
		Filter("user_id", user_id).
		Filter("domain", domain).
		Update(orm.Params{
		"link": link,
	})
	if err != nil {
		u.MsgBack("更新阅读进度失败", 0)
	}
	u.MsgBack("更新阅读进度成功", 1)
}

//退出登录
func (u *UserController) Logout() {
	u.DelSession("user")
	u.DelSession("user_id")
	u.MsgBack("退出成功!", 1)
}

//删除书架图书
func (u *UserController) DelBook() {
	user_id := u.Data["user_id"].(int)
	domain := u.GetString("domain")
	_, err := u.o.QueryTable(new(models.Bookshelf).TableName()).
		Filter("user_id", user_id).
		Filter("domain", domain).
		Delete()
	if err != nil {
		u.MsgBack("删除失败", 0)
	}
	u.MsgBack("删除成功", 1)
}
