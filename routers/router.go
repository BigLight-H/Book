package routers

import (
	"fiction_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/home/:name", &controllers.MainController{},"get:Home")
	beego.Router("/book/list", &controllers.MainController{},"get:List")
	beego.Router("/book/detail", &controllers.MainController{},"get:Detail")
}
