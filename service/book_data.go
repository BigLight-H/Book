package service

import (
	"fiction_web/models"
	"github.com/gocolly/colly"
	"strconv"
)

//书目录
func BookList(list models.Lists, id int, links string) []map[string]string {
	c := colly.NewCollector()
	info := []map[string]string{}
	c.OnXML(list.Root, func(e *colly.XMLElement) {
		link := ""
		switch list.Id {
			case 2:
				link = links + e.ChildText(list.Link)
				break
			case 3:
				link = list.Domain + e.ChildText(list.Link)
				break
			default:
				link = e.ChildText(list.Link)
				break
		}
		name := e.ChildText(list.Name)
		id := strconv.Itoa(id)
		info = append(
			info,
			map[string]string{"link": link, "name": name, "id": id})
	})
	c.Visit(links)
	return info
}

//书正文
func BookContent(con models.Content, links string) []map[string]string {
	c := colly.NewCollector()
	info := []map[string]string{}
	c.OnXML(con.Root, func(e *colly.XMLElement) {
		name := e.ChildText(con.Name)
		var ele  = NewXMLElement(e)
		content := ele.ChildHtml(con.Content)
		s_page := ""
		x_page := ""
		list := ""
		switch con.Id {
			case 3:
				s_page = con.Domain + e.ChildText(con.SPage)
				x_page = con.Domain + e.ChildText(con.XPage)
				list = con.Domain + e.ChildText(con.List)
				break
			default:
				s_page = e.ChildText(con.SPage)
				x_page = e.ChildText(con.XPage)
				list = e.ChildText(con.List)
				break
		}
		info = append(
			info,
			map[string]string{"name": name, "content": content, "s_page": s_page, "x_page": x_page, "list": list})
	})
	c.Visit(links)
	return info
}
//书图片,简介
func BookSynosis(con models.Synopsis, links string) []map[string]string  {
	c := colly.NewCollector()
	info := []map[string]string{}
	c.OnXML(con.Root, func(e *colly.XMLElement) {
		name := e.ChildText(con.Name)
		writer := e.ChildText(con.Writer)
		img := e.ChildText(con.Img)
		synopsis := e.ChildText(con.Synopsis)
		renew_time := e.ChildText(con.RenewTime)
		info = append(
			info,
			map[string]string{"name": name, "writer": writer, "img": img, "synopsis": synopsis, "renew_time": renew_time})
	})
	c.Visit(links)
	return info
}

//检测书本是否更新
func BookSynosisCheck(con models.Synopsis, links string) string  {
	c := colly.NewCollector()
	info := ""
	c.OnXML(con.Root, func(e *colly.XMLElement) {
		renew_time := e.ChildText(con.RenewTime)
		info = renew_time
	})
	c.Visit(links)
	return info
}