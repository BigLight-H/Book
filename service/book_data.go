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
		content := e.ChildText(con.Content)
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