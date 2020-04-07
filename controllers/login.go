package controllers

import (
	"fiction_web/models"
	"fiction_web/util"
)

type LoginController struct {
	BaseController
}

//注册
func (p *LoginController) Register() {
	name := p.GetString("name")
	email := p.GetString("email")
	pwd := p.GetString("pwd")
	if util.VerifyEmailFormat(email)  != true {
		p.MsgBack("邮箱格式不正确!", 0)
		panic("go back!")
	}
	if p.checkUserName(name) != true {
		p.MsgBack("用户名已经存在!", 0)
		panic("go back!")
	}
	if p.checkEmail(email) != true {
		p.MsgBack("邮箱已存在!", 0)
		panic("go back!")
	}
	user := models.Users{}
	user.Name = name
	user.Email = email
	user.Pwd = util.Md5(pwd)
	_, err := p.o.Insert(&user)
	if err != nil {
		p.MsgBack("注册失败!", 0)
		panic("go back!")
	}
	p.MsgBack("注册成功!",1)
}

//登录
func (p *LoginController) Login() {
	name := p.GetString("name")
	pwd := util.Md5(p.GetString("pwd"))
	if name == "" {
		p.MsgBack("用户名不能为空!", 0)
		panic("用户名不能为空!")
	}
	user := []*models.Users{}
	err := p.o.QueryTable(new(models.Users).TableName()).Filter("name", name).Filter("pwd", pwd).One(&user)
	if err != nil {
		p.MsgBack("登录失败!", 0)
	} else {
		token := util.GenerateToken(86400, user[0].Id, user[0].Name)//token有效期时长一天
		p.MsgBack(token, 1)
	}
}

//验证用户名是否存在
func (p *LoginController) checkUserName(name string) bool {
	num,err := p.o.QueryTable(new(models.Users).TableName()).Filter("name", name).Count()
	if err ==nil {
		if num == 0 {
			return true
		}
	}
	return false
}

//验证邮箱是否存在
func (p *LoginController) checkEmail(email string) bool {
	num,err := p.o.QueryTable(new(models.Users).TableName()).Filter("email", email).Count()
	if err ==nil {
		if num == 0 {
			return true
		}
	}
	return false
}