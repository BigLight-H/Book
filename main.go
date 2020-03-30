package main

import (
	"fiction_web/models"
	_ "fiction_web/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	orm.Debug = true
	models.Init()
}

func main() {
	beego.Run()
}


