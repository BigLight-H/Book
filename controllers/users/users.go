package users

import (
	"fiction_web/controllers"
	"fiction_web/models"
	"fiction_web/util"
	"github.com/astaxie/beego/orm"
	"strings"
)

type UserController struct {
	controllers.BaseController
}

type bookshelf struct {
	Id    int         `form:"-"`
	HubId     int     `form:"id"`
	Link      string  `form:"link"`
	Img       string  `form:"img"`
	Domain    string  `form:"domain"`
	BookName  string  `form:"name"`
	Author    string  `form:"writer"`
	RenewTime string  `form:"renew_time"`
}

//添加书架
func (u *UserController) AddBookshelf() {
	token := strings.Fields(u.Ctx.Input.Header("Authorization"))//获取token
	b := bookshelf{}
	u.ParseForm(&b)
	hub_id := b.HubId
	link := b.Link
	domain := b.Domain
	img := b.Img
	author := b.Author
	book_name := b.BookName
	user_id := util.GetTokenUserId(token[1])
	renew_time := b.RenewTime
	u.VerificationBook(user_id, domain)
	o := orm.NewOrm()
	book := models.Bookshelf{}
	book.Link = link
	book.UserId = user_id
	book.HubId = hub_id
	book.Domain = domain
	book.BookName = book_name
	book.Img = img
	book.Author = author
	book.RenewTime = renew_time
	o.Insert(&book)
	u.MsgBack("添加书签成功", 1)
}

//查看书架
func (u *UserController) GetBookshelf() {
	o := orm.NewOrm()
	user_id := u.Data["user_id"].(int)
	book := []models.Bookshelf{}
	o.QueryTable(new(models.Bookshelf).TableName()).Filter("user_id", user_id).All(&book)
	u.Data["json"] = book
	u.ServeJSON()
}

//查看该书本是否已经加入书架
func (u *UserController) VerificationBook(id int, domain string) {
	o := orm.NewOrm()
	num,_ := o.QueryTable(new(models.Bookshelf).TableName()).
		Filter("user_id", id).Filter("domain", domain).
		Count()
	if num > 0 {
		u.MsgBack("书本已存在", 0)
	}
}

//查看该书本是否已经加入书架
func (u *UserController) VerificationBooks() {
	token := strings.Fields(u.Ctx.Input.Header("Authorization"))//获取token
	id :=  util.GetTokenUserId(token[1])
	domain := u.GetString("domain")
	o := orm.NewOrm()
	num,_ := o.QueryTable(new(models.Bookshelf).TableName()).
		Filter("user_id", id).Filter("domain", domain).
		Count()
	if num > 0 {
		u.MsgBack("书本已存在", 0)
	} else {
		u.MsgBack("书本不存在", 1)
	}
}

//更新图书阅读进度
func (u *UserController) UpdateBookSchedule() {
	o := orm.NewOrm()
	user_id := u.Data["user_id"].(int)
	domain := u.GetString("domain")
	link := u.GetString("link")
	_, err := o.QueryTable(new(models.Bookshelf).TableName()).
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
	o := orm.NewOrm()
	user_id := u.Data["user_id"].(int)
	domain := u.GetString("domain")
	_, err := o.QueryTable(new(models.Bookshelf).TableName()).
		Filter("user_id", user_id).
		Filter("domain", domain).
		Delete()
	if err != nil {
		u.MsgBack("删除失败", 0)
	}
	u.MsgBack("删除成功", 1)
}
