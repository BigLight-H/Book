package users

import (
	"fiction_web/controllers"
	"fiction_web/models"
	"fiction_web/service"
	"fiction_web/util"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"strings"
)

type UserController struct {
	controllers.BaseController
}

type bookshelf struct {
	Id        int     `form:"-"`
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
	token := strings.Fields(u.Ctx.Input.Header("Authorization"))//获取token
	user_id := util.GetTokenUserId(token[1])
	book := []models.Bookshelf{}
	o.QueryTable(new(models.Bookshelf).TableName()).Filter("user_id", user_id).All(&book)
	info := []map[string]string{}
	for _, value := range book{
		new_renew := u.CheckNews(value.HubId, value.Domain, value.RenewTime)
		info = append(
			info,
			map[string]string{
				"UserId": strconv.Itoa(value.UserId),
				"HubId": strconv.Itoa(value.HubId),
				"Link": value.Link,
				"Img": value.Img,
				"Domain": value.Domain,
				"BookName": value.BookName,
				"Author": value.Author,
				"RenewTime": value.RenewTime,
				"Status":new_renew["status"],
				"NewRenew":new_renew["new_renew_time"],
			})
	}
	u.Data["json"] = info
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
		panic("书本已存在")
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
	token := strings.Fields(u.Ctx.Input.Header("Authorization"))//获取token
	o := orm.NewOrm()
	user_id := util.GetTokenUserId(token[1])
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

//删除书架图书
func (u *UserController) DelBook() {
	token := strings.Fields(u.Ctx.Input.Header("Authorization"))//获取token
	o := orm.NewOrm()
	user_id := util.GetTokenUserId(token[1])
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

//检测书本是否有更新
func (u *UserController) CheckNews(hubId int, link string, renewTime string) map[string]string {
	o := orm.NewOrm()
	con := models.Synopsis{Id:hubId}
	o.Read(&con)
	str := service.BookSynosisCheck(con, link)
	spew.Dump(renewTime, str)
	info := map[string]string{}
	info["status"] = strconv.Itoa(strings.Compare(renewTime, str))
	info["new_renew_time"] = str
	return info
}

//更新最新章节tag
func (u *UserController) ChangeRenewTime()  {
	token := strings.Fields(u.Ctx.Input.Header("Authorization"))//获取token
	o := orm.NewOrm()
	user_id := util.GetTokenUserId(token[1])
	domain := u.GetString("domain")
	new_renew := u.GetString("new_renew")
	_, err := o.QueryTable(new(models.Bookshelf).TableName()).
		Filter("user_id", user_id).
		Filter("domain", domain).
		Update(orm.Params{
			"renew_time": new_renew,
		})
	if err != nil {
		u.MsgBack("更新最新章节tag失败", 0)
	}
	u.MsgBack("更新最新章节tag成功", 1)
}
