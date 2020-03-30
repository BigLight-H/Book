package service

import (
	"fiction_web/models"
	"fiction_web/util"
	"github.com/astaxie/beego/orm"
	"github.com/gocolly/colly"
)

//酷小说
func KuXiaoShuo(s string) []map[string]string {
	o := orm.NewOrm()
	c := colly.NewCollector()
	hub := models.Hub{Id:1}
	o.Read(&hub)
	info := []map[string]string{}
	c.OnXML(hub.Root, func(e *colly.XMLElement) {
		link := e.ChildText(hub.Link)
		name := e.ChildText(hub.Name)
		uname := e.ChildText(hub.Author)
		id := "1"
		info = util.BackInfoMap(info, link, name, uname, id)
	})
	c.Visit(hub.Suffix + s)
	return info
}



//笔趣阁
func BiQuGe(s string) []map[string]string {
	c := colly.NewCollector()
	str := "https://www.biquge.info"
	info := []map[string]string{}
	c.OnXML("//tbody/tr[1]/following-sibling::tr", func(e *colly.XMLElement) {
		link := str + e.ChildText("./td[1]/a/@href")
		name := e.ChildText("./td[1]/a")
		uname := e.ChildText("./td[3]")
		id := "2"
		info = util.BackInfoMap(info, link, name, uname, id)
	})
	c.Visit(str + "/modules/article/search.php?searchkey=" + s)
	return info
}

//笔趣阁VIP
func VipZw(s string) []map[string]string  {
	c := colly.NewCollector()
	str := "http://www.vipzw.com/"
	info := []map[string]string{}
	c.OnXML("//div[@class=\"item\"]", func(e *colly.XMLElement) {
		link := e.ChildText("./dl/dt/a/@href")
		name := e.ChildText("./dl/dt/a")
		uname := e.ChildText("./dl/dt/span")
		id := "3"
		info = util.BackInfoMap(info, link, name, uname, id)
	})
	c.Visit(str + "/search.php?searchkey=" + s)
	return info
}

//合并数据
func MergeMap(s string) []map[string]string {
	info := KuXiaoShuo(s)
	biqu := BiQuGe(s)
	vipbiqu := VipZw(s)
	return MergeData(vipbiqu, MergeData(biqu, info))
}

//合并数据函数
func MergeData(data []map[string]string, info []map[string]string) []map[string]string {
	for _, v := range data {
		name := v["name"]
		link := v["link"]
		uname := v["uname"]
		id := v["id"]
		info = util.BackInfoMap(info, link, name, uname, id)
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
