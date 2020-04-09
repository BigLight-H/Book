package controllers

import (
	"encoding/json"
	"fiction_web/models"
	"fiction_web/util"
	"github.com/astaxie/beego"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
)

type LoginController struct {
	BaseController
}

type TenXunRequest struct {
	Response     string     `json:"response"`
	Evil_level   string     `json:"evil_level"`
	Err_msg      string     `json:"err_msg"`
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
	ticket := p.GetString("ticket")
	randstr := p.GetString("randstr")
	userIp := util.RemoteIp(p.Ctx.Request)
	pwd := util.Md5(p.GetString("pwd"))
	v := p.httpGet(ticket, randstr, userIp)
	spew.Dump(v)
	if v < 1 {
		p.MsgBack("验证码验证失败4	", 0)
		panic("验证码验证失败")
	}
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
//腾旭防水墙验证
func (p *LoginController) httpGet(Ticket string, Randstr string, UserIP string) int {
	resp, err := http.Get("https://ssl.captcha.qq.com/ticket/verify?aid="+beego.AppConfig.String("aid")+"&AppSecretKey="+beego.AppConfig.String("AppSecretKey")+"&Ticket="+Ticket+"&Randstr="+Randstr+"&UserIP="+UserIP)
	if err != nil {
		p.MsgBack("验证码验证失败1", 0)
		panic("验证码验证失败")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.MsgBack("验证码验证失败2", 0)
		panic("验证码验证失败")
	}
	var a TenXunRequest
	json.Unmarshal(body, &a)
	return util.StrToInt(a.Response)
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