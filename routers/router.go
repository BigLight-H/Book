package routers

import (
	"fiction_web/controllers"
	"fiction_web/controllers/users"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //搜索图书
	beego.Router("/home/:name", &controllers.MainController{},"get:Home")
    //获取图书目录列表
	beego.Router("/book/list", &controllers.MainController{},"get:List")
    //获取图书内容详情
	beego.Router("/book/detail", &controllers.MainController{},"get:Detail")
    //文章目录介绍部分
	beego.Router("/book/synopsis", &controllers.MainController{},"get:Synopsis")

    //登录/注册
	beego.Router("/register", &controllers.LoginController{},"post:Register")
	beego.Router("/login", &controllers.LoginController{},"post:Login")

    //用户登录后
    //获取书架数据
	beego.Router("/user/books", &users.UserController{},"get:GetBookshelf")
    //添加到书架
	beego.Router("/user/add/books", &users.UserController{},"post:AddBookshelf")
    //查看书本是否已经存在书架
	beego.Router("/user/verification/books", &users.UserController{},"get:VerificationBooks")
	//更新图书阅读进度
	beego.Router("/user/books/update", &users.UserController{},"post:UpdateBookSchedule")
	//更新图书最新章节tag
	beego.Router("/user/books/renew/time", &users.UserController{},"post:ChangeRenewTime")
	//删除图书
	beego.Router("/user/books/del", &users.UserController{},"post:DelBook")

    //书源首页
	beego.Router("/book/source", &controllers.BookSourceController{},"get:BookSource")
    //分类链接,名
	beego.Router("/book/source/types", &controllers.BookSourceController{},"get:BookTypes")
    //分类列表
	beego.Router("/book/source/type", &controllers.BookSourceController{},"get:BookType")
    //排行列表
	beego.Router("/book/source/board/list", &controllers.BookSourceController{},"get:LeaderBoard")
    //排行链接,名
	beego.Router("/book/source/board", &controllers.BookSourceController{},"get:Leader")
    //获取完本数据列表
	beego.Router("/book/source/book/end", &controllers.BookSourceController{},"get:BookEnd")
}
