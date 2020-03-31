package service

import (
	"fiction_web/models"
	"fiction_web/util"
	"github.com/astaxie/beego/orm"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func GetBook(s string) []map[string]string {
	o := orm.NewOrm()
	c := colly.NewCollector()
	info := []map[string]string{}
	hub := []models.Hub{}
	o.QueryTable(new(models.Hub).TableName()).All(&hub)
	for _, v := range hub {
		c.OnXML(v.Root, func(e *colly.XMLElement) {
			link := ""
			if strings.Contains(e.ChildText(v.Link), "http") == false {
				link = v.BookHub + e.ChildText(v.Link)
			} else {
				link = e.ChildText(v.Link)
			}
			name := e.ChildText(v.Name)
			uname := e.ChildText(v.Author)
			id := strconv.Itoa(v.Id)
			info = util.BackInfoMap(info, link, name, uname, id)
		})
		c.Visit(v.Suffix + s)
	}
	return Deduplication(info)
}

//数据去重
func Deduplication(data []map[string]string) []map[string]string {
	src := make(map[string]interface{})
	info := []map[string]string{}
	for _, v := range data {
		if _,ok:=src[v["name"]];ok {
			continue
		} else {
			src[v["name"]] = v["name"]
			name := v["name"]
			link := v["link"]
			uname := v["uname"]
			id := v["id"]
			info = util.BackInfoMap(info, link, name, uname, id)
		}
	}
	return info
}
